package service

import (
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/repository"
)

type VideoCreatePostFlow struct {
	// 请求需要的数据
	Title    string `json:"title"`
	PlayURL  string `json:"play_url"`
	CoverURL string `json:"cover_url"`
	UserId   int64  `json:"user_id"`
	// 响应需要的数据
	VideoId int64 `json:"video_id"`
}

func NewVideoCreatePostFlow(title, playUrl, coverUrl string, userId int64) *VideoCreatePostFlow {
	return &VideoCreatePostFlow{
		Title:    title,
		PlayURL:  playUrl,
		CoverURL: coverUrl,
		UserId:   userId,
	}
}
func VideoCreatePost(title, playUrl, coverUrl string, userId int64) (int64, errno.ErrNo) {
	return NewVideoCreatePostFlow(title, playUrl, coverUrl, userId).Do()
}

func (f *VideoCreatePostFlow) Do() (int64, errno.ErrNo) {
	// 1. check params
	if err := f.checkParam(); err.ErrCode != 0 {
		return 0, err
	}
	// 2. create video information
	if err := f.create(); err.ErrCode != 0 {
		return 0, err
	}
	return f.VideoId, errno.Success
}
func (f *VideoCreatePostFlow) checkParam() errno.ErrNo {
	if f.Title == "" || f.PlayURL == "" || f.CoverURL == "" || f.UserId == 0 {
		return errno.ParamErr
	}
	return errno.Success
}
func (f *VideoCreatePostFlow) create() errno.ErrNo {
	video := &repository.Video{
		Title:    f.Title,
		PlayURL:  f.PlayURL,
		CoverURL: f.CoverURL,
		UserId:   f.UserId,
	}
	videoId, err := repository.NewVideoDaoInstance().CreateVideo(video)
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.VideoId = videoId
	return errno.Success
}
