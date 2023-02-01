package service

import (
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/repository"
)

type RelationshipQueryTListGetFlow struct {
	// 请求需要的数据
	UserId int64 `json:"user_id"` // 查询UserId的舔狗列表
	// 响应需要的数据
	SubscribeUsers []*repository.User `json:"subscribe_users"`
}

func NewRelationshipQueryTListGetFlow(userId int64) *RelationshipQueryTListGetFlow {
	return &RelationshipQueryTListGetFlow{UserId: userId}
}
func RelationshipQueryTListGet(userId int64) (*RelationshipQueryTListGetFlow, errno.ErrNo) {
	return NewRelationshipQueryTListGetFlow(userId).Do()
}

func (f *RelationshipQueryTListGetFlow) Do() (*RelationshipQueryTListGetFlow, errno.ErrNo) {
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
func (f *RelationshipQueryTListGetFlow) checkParam() errno.ErrNo {
	if f.UserId == 0 {
		return errno.ParamErr
	}
	return errno.Success
}
func (f *RelationshipQueryTListGetFlow) queryUser() errno.ErrNo {
	rs, err := repository.NewRelationshipDaoInstance().QueryRelationshipByUserId(f.UserId)
	if err != nil {
		return errno.ConvertErr(err)
	}
	subscribeUserIds := make([]int64, 0)
	for _, r := range rs {
		subscribeUserIds = append(subscribeUserIds, r.ID)
	}
	subscribeUsers, err := repository.NewUserDaoInstance().MQueryUserById(subscribeUserIds)
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.SubscribeUsers = subscribeUsers
	return errno.Success
}
