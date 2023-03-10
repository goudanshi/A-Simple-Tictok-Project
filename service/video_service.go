package service

import (
	"bytes"
	"douyin/repository"
	"douyin/util"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"sync"
	"time"
)

type VideoUser struct {
	repository.Video
	IsFavorite bool `json:"is_favorite"`
	User       User `json:"author"`
}

func GetVideoFeed(name string) (io.Reader, int64, time.Time, error) {
	return util.GetObjectWithSize(name)
}

func GetVideoCover(name string) (io.Reader, int64, time.Time, error) {
	return util.GetObjectWithSize(name)
}

func PublishVideo(data *multipart.FileHeader, title string, userId int64) (int64, error) {
	file, _ := data.Open()
	defer file.Close()

	fileName := strconv.FormatInt(userId, 10) + "/" + strconv.FormatInt(time.Now().Unix(), 10)

	err := util.PutVideo(fileName, file, data.Size)
	if err != nil {
		return -1, err
	}
	reader := bytes.NewBuffer(nil)
	err = ffmpeg.Input("http://localhost:8100"+util.GET_VIDEO_PATH+"?name="+fileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 5)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return -1, err
	}
	coverName := fileName + "-cover.jpeg"
	util.PutJpg(coverName, reader, int64(reader.Len()))

	return repository.GetVideoDaoInstance().Add(&repository.Video{
		PublisherId:   userId,
		Title:         title,
		VideoUrl:      util.OSS_SHARE_HOST + fileName,
		CoverUrl:      util.OSS_SHARE_HOST + coverName,
		FavoriteCount: 0,
		CommentCount:  0,
		CreateDate:    time.Now(),
	})
}

func VideoList(userId int64) ([]VideoUser, error) {
	videoData, err := repository.GetVideoDaoInstance().Query(util.FEED_LIMIT)
	if err != nil {
		return nil, err
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	var followMap, favoriteMap map[int64]struct{}
	var followErr, favoriteErr error
	earlyDate := videoData[len(videoData)-1].Video.CreateDate
	go func() {
		followMap, followErr = getFollowMap(userId)
		waitGroup.Done()
	}()
	go func() {
		favoriteMap, favoriteErr = getFavoriteMap(userId, earlyDate)
		waitGroup.Done()
	}()
	waitGroup.Wait()
	if followErr != nil {
		return nil, followErr
	}
	if favoriteErr != nil {
		return nil, favoriteErr
	}

	result := make([]VideoUser, len(videoData))
	for i, item := range videoData {
		raw := VideoUser{item.Video, true, *convertUser(&item.User, true)}
		if _, ok := followMap[item.User.Id]; !ok {
			raw.User.IsFollow = false
		}
		if _, ok := favoriteMap[item.Video.Id]; !ok {
			raw.IsFavorite = false
		}
		result[i] = raw
	}

	return result, nil
}

func QueryUserVideo(userId int64) ([]VideoUser, error) {
	data, err := repository.GetVideoDaoInstance().QueryUserVideo(userId, 30)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []VideoUser{}, err
	}
	favoriteMap, err := getFavoriteMap(userId, data[len(data)-1].Video.CreateDate)
	if err != nil {
		return nil, err
	}
	result := make([]VideoUser, len(data))
	for i, item := range data {
		isFavorite := false
		if _, ok := favoriteMap[item.Video.Id]; ok {
			isFavorite = true
		}
		result[i] = VideoUser{
			Video:      item.Video,
			IsFavorite: isFavorite,
			User:       *convertUser(&item.User, false),
		}
	}

	return result, nil
}
