package service

import (
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/repository"
)

type VideoQueryGetFlow struct {
	// 请求需要的数据
	UserId int64 `json:"user_id"`
	// 响应需要的数据
	Author *repository.User    `json:"author"`
	Videos []*repository.Video `json:"videos"`
}

func NewVideoQueryGetFlow(userId int64) *VideoQueryGetFlow {
	return &VideoQueryGetFlow{UserId: userId}
}

func VideoQueryGet(userId int64) (*VideoQueryGetFlow, errno.ErrNo) {
	return NewVideoQueryGetFlow(userId).Do()
}
func (f *VideoQueryGetFlow) Do() (*VideoQueryGetFlow, errno.ErrNo) {
	// 1. check param
	if err := f.checkParam(); err.ErrCode != 0 {
		return nil, err
	}
	// 2. query user information
	if err := f.queryUser(); err.ErrCode != 0 {
		return nil, err
	}
	// 3. query video information
	if err := f.queryVideo(); err.ErrCode != 0 {
		return nil, err
	}
	return f, errno.Success
}
func (f *VideoQueryGetFlow) checkParam() errno.ErrNo {
	if f.UserId == 0 {
		return errno.ParamErr
	}
	return errno.Success
}

func (f *VideoQueryGetFlow) queryUser() errno.ErrNo {
	users, err := repository.NewUserDaoInstance().MQueryUserById([]int64{f.UserId})
	if err != nil {
		return errno.ConvertErr(err)
	}
	if len(users) == 0 {
		return errno.UserNotExistErr
	}
	f.Author = users[0]
	return errno.Success
}

func (f *VideoQueryGetFlow) queryVideo() errno.ErrNo {
	videos, err := repository.NewVideoDaoInstance().QueryVideoByUserId(f.UserId)
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.Videos = videos
	return errno.Success
}
