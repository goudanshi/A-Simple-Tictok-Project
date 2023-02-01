package service

import (
	"crypto/md5"
	"fmt"
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/repository"
	"io"
)

func UserCheckPost(username, password string) (*repository.User, errno.ErrNo) {
	return NewUserCheckPost(username, password).Do()
}
func NewUserCheckPost(username, password string) *UserCheckPostFlow {
	return &UserCheckPostFlow{
		Username: username,
		Password: password,
	}
}

type UserCheckPostFlow struct {
	// 请求需要的数据
	Username string `json:"username"`
	Password string `json:"password"`
	// 响应需要的数据
	User *repository.User `json:"user"`
}

func (f *UserCheckPostFlow) Do() (*repository.User, errno.ErrNo) {
	// checkParam
	if err := f.checkParam(); err.ErrCode != 0 {
		return nil, err
	}
	// queryUser
	if err := f.queryUser(); err.ErrCode != 0 {
		return nil, err
	}
	// decryptPassword
	if err := f.decryptPassword(); err.ErrCode != 0 {
		return nil, err
	}
	return f.User, errno.Success
}

func (f *UserCheckPostFlow) checkParam() errno.ErrNo {
	if len(f.Username) == 0 || len(f.Password) == 0 {
		return errno.ParamErr
	}
	return errno.Success
}

func (f *UserCheckPostFlow) queryUser() errno.ErrNo {
	user, err := repository.NewUserDaoInstance().QueryUserByUsername(f.Username)
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.User = user
	return errno.Success
}

func (f *UserCheckPostFlow) decryptPassword() errno.ErrNo {
	h := md5.New()
	_, err := io.WriteString(h, f.Password) // f.Password 明文的密码
	if err != nil {
		return errno.ServiceErr
	}
	password := fmt.Sprintf("%x", h.Sum(nil)) // 加密后的密码
	if password != f.User.Password {          // f.User.Password 数据库中存储的加密的密码
		return errno.AuthorizationFailedErr
	}
	return errno.Success
}
