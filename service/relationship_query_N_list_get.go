package service

import (
	"github.com/A-Simple-Tictok-Project/douyin/errno"
	"github.com/A-Simple-Tictok-Project/douyin/repository"
)

type RelationshipQueryNListGetFlow struct {
	// 请求需要的数据
	SubscribeId int64 `json:"subscribe_id"` // 查询SubscribeId的女神列表
	// 响应需要的数据
	Users []*repository.User `json:"users"`
}

func NewRelationshipQueryNListGetFlow(subscribeId int64) *RelationshipQueryNListGetFlow {
	return &RelationshipQueryNListGetFlow{SubscribeId: subscribeId}
}
func RelationshipQueryNListGet(userId int64) (*RelationshipQueryNListGetFlow, errno.ErrNo) {
	return NewRelationshipQueryNListGetFlow(userId).Do()
}

func (f *RelationshipQueryNListGetFlow) Do() (*RelationshipQueryNListGetFlow, errno.ErrNo) {
	// 1. check param
	if err := f.checkParam(); err.ErrCode != 0 {
		return nil, err
	}
	// 2. query user information
	if err := f.queryUser(); err.ErrCode != 0 {
		return nil, err
	}
	return f, errno.Success
}
func (f *RelationshipQueryNListGetFlow) checkParam() errno.ErrNo {
	if f.SubscribeId == 0 {
		return errno.ParamErr
	}
	return errno.Success
}
func (f *RelationshipQueryNListGetFlow) queryUser() errno.ErrNo {
	rs, err := repository.NewRelationshipDaoInstance().QueryRelationshipByUserId(f.SubscribeId)
	if err != nil {
		return errno.ConvertErr(err)
	}
	userIds := make([]int64, 0)
	for _, r := range rs {
		userIds = append(userIds, r.ID)
	}
	users, err := repository.NewUserDaoInstance().MQueryUserById(userIds)
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.Users = users
	return errno.Success
}
