package service

import (
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/repository"
)

type LikeQueryListGetFlow struct {
	// 请求需要的数据
	UserId int64 `json:"user_id"`
	// 响应需要的数据
	Videos    []*repository.Video        `json:"videos"`
	AuthorMap map[int64]*repository.User `json:"author_map"`
}

func NewLikeQueryListGet(userId int64) *LikeQueryListGetFlow {
	return &LikeQueryListGetFlow{UserId: userId}
}

func LikeQueryList(userId int64) (*LikeQueryListGetFlow, errno.ErrNo) {
	return NewLikeQueryListGet(userId).Do()
}

func (f *LikeQueryListGetFlow) Do() (*LikeQueryListGetFlow, errno.ErrNo) {
	// 1. check param
	if err := f.checkParam(); err.ErrCode != 0 {
		return nil, err
	}
	// 2. query like video list information
	if err := f.queryVideo(); err.ErrCode != 0 {
		return nil, err
	}
	// 3. query every video author information
	if err := f.queryUser(); err.ErrCode != 0 {
		return nil, err
	}
	return f, errno.Success
}

func (f *LikeQueryListGetFlow) checkParam() errno.ErrNo {
	if f.UserId == 0 {
		return errno.ParamErr
	}
	return errno.Success
}

func (f *LikeQueryListGetFlow) queryVideo() errno.ErrNo {
	likes, err := repository.NewLikeDaoInstance().QueryLikeByUserId(f.UserId)
	if err != nil {
		return errno.ConvertErr(err)
	}
	videoIds := make([]int64, 0)
	for _, like := range likes {
		videoIds = append(videoIds, like.VideoId)
	}
	videos, err := repository.NewVideoDaoInstance().QueryVideoByVideoId(videoIds)
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.Videos = videos
	return errno.Success
}

func (f *LikeQueryListGetFlow) queryUser() errno.ErrNo {
	userIds := make([]int64, 0)
	authorMap := make(map[int64]*repository.User)
	for _, video := range f.Videos {
		userIds = append(userIds, video.UserId)
	}
	users, err := repository.NewUserDaoInstance().MQueryUserById(userIds)
	if err != nil {
		return errno.ConvertErr(err)
	}
	for _, user := range users {
		authorMap[user.ID] = user
	}
	f.AuthorMap = authorMap
	return errno.Success
}
