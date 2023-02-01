package service

import (
	"github.com/A-Simple-Tictok-Project/douyin/errno"
	"github.com/A-Simple-Tictok-Project/douyin/repository"
)

type UserQueryGetFlow struct {
	// 请求需要的数据
	UserId int64 `json:"user_id"`
	// 响应需要的数据
	Users []*repository.User `json:"users"`
}

func NewUserQueryGetFlow(userId int64) *UserQueryGetFlow {
	return &UserQueryGetFlow{UserId: userId}
}

func UserQueryGet(userId int64) ([]*repository.User, errno.ErrNo) {
	return NewUserQueryGetFlow(userId).Do()
}

func (f *UserQueryGetFlow) Do() ([]*repository.User, errno.ErrNo) {
	// 1. check params
	if err := f.checkParam(); err.ErrCode != 0 {
		return nil, err
	}
	// 2. query user information
	if err := f.queryUser(); err.ErrCode != 0 {
		return nil, err
	}
	return f.Users, errno.Success
}
func (f *UserQueryGetFlow) checkParam() errno.ErrNo {
	if f.UserId == 0 {
		return errno.ParamErr
	}
	return errno.Success
}
func (f *UserQueryGetFlow) queryUser() errno.ErrNo {
	users, err := repository.NewUserDaoInstance().MQueryUserById([]int64{f.UserId})
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.Users = users
	return errno.Success
}
