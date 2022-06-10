package service

import (
	"douyin/repository"
	"douyin/util"
	"sync"
	"time"
)

func NewFavorite(userId int64, videoId int64) (int64, error) {
	id, err := repository.GetFavoriteDaoInstance().Add(&repository.Favorite{
		UserId:     userId,
		VideoId:    videoId,
		CreateDate: time.Now(),
	})
	if err != nil {
		return -1, err
	}

	go func(videoId int64) {
		videoDao := repository.GetVideoDaoInstance()
		video, e := videoDao.QueryById(videoId)
		if e != nil {
			util.Logger.Error(e.Error())
			return
		}
		video.FavoriteCount++
		videoDao.Update(video)
	}(videoId)
	return id, err
}

func CancelFavorite(userId int64, videoId int64) error {
	err := repository.GetFavoriteDaoInstance().DeleteByUserVideo(userId, videoId)
	if err != nil {
		return err
	}
	go func(videoId int64) {
		videoDao := repository.GetVideoDaoInstance()
		video, er := videoDao.QueryById(videoId)
		if er != nil {
			util.Logger.Error(er.Error())
			return
		}
		video.FavoriteCount--
		er = videoDao.Update(video)
		if err != nil {
			util.Logger.Error(er.Error())
		}
	}(videoId)
	return nil
}

func QueryFavoriteVideo(userId int64) ([]VideoUser, error) {
	var data []repository.VideoUser
	var followMap map[int64]struct{}
	var dataErr, followErr error
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		data, dataErr = repository.GetFavoriteDaoInstance().QueryFavoriteVideo(userId)
		waitGroup.Done()
	}()
	go func() {
		followMap, followErr = getFollowMap(userId)
		waitGroup.Done()
	}()
	waitGroup.Wait()

	if dataErr != nil {
		return nil, dataErr
	}
	if followErr != nil {
		return nil, followErr
	}

	result := make([]VideoUser, len(data))
	for i, item := range data {
		var isFollow = false
		if _, ok := followMap[item.User.Id]; ok {
			isFollow = true
		}
		result[i] = VideoUser{
			Video:      item.Video,
			IsFavorite: true,
			User:       *convertUser(&item.User, isFollow),
		}
	}

	return result, nil
}

func getFavoriteMap(userId int64, earlyTime time.Time) (map[int64]struct{}, error) {
	data, err := repository.GetFavoriteDaoInstance().QueryVideoByUserAndEarlyDate(userId, earlyTime)
	if err != nil {
		return nil, err
	}
	favoriteMap := make(map[int64]struct{})
	for _, item := range data {
		favoriteMap[item] = struct{}{}
	}

	return favoriteMap, nil
}
