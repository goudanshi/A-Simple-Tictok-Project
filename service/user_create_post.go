package service

import (
	"crypto/md5"
	"fmt"
	"github.com/A-Simple-Tictok-Project/douyin/errno"
	"github.com/A-Simple-Tictok-Project/douyin/repository"
	"io"
)

func UserCreatePost(username, password string) (int64, errno.ErrNo) {
	return NewUserCreatePost(username, password).Do()
}

func NewUserCreatePost(username, password string) *UserCreatePostFlow {
	return &UserCreatePostFlow{
		Username: username,
		Password: password,
	}
}

type UserCreatePostFlow struct {
	// 请求需要的数据
	Username string
	Password string

	// 响应需要的数据
	UserId int64
}

func (f *UserCreatePostFlow) Do() (int64, errno.ErrNo) {
	// 1. check params
	if err := f.checkParam(); err.ErrCode != 0 {
		return 0, err
	}
	// 2. crypto password
	if err := f.cryptPassword(); err.ErrCode != 0 {
		return 0, err
	}
	// 3. create user info
	if err := f.create(); err.ErrCode != 0 {
		return 0, err
	}

	return f.UserId, errno.Success
}

func (f *UserCreatePostFlow) checkParam() errno.ErrNo {
	if len(f.Username) == 0 || len(f.Password) == 0 {
		return errno.ParamErr
	}
	return errno.Success
}

func (f *UserCreatePostFlow) cryptPassword() errno.ErrNo {
	h := md5.New()
	_, err := io.WriteString(h, f.Password)
	if err != nil {
		return errno.ServiceErr
	}
	f.Password = fmt.Sprintf("%x", h.Sum(nil))
	return errno.Success
}

func (f *UserCreatePostFlow) create() errno.ErrNo {
	user := &repository.User{
		Username: f.Username,
		Password: f.Password,
	}
	if err := repository.NewUserDaoInstance().CreateUser(user); err != nil {
		return errno.ConvertErr(err)
	}
	f.UserId = user.ID
	return errno.Success
}
