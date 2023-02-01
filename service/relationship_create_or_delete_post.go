package service

import (
	"github.com/A-Simple-Tictok-Project/douyin/constants"
	"github.com/A-Simple-Tictok-Project/douyin/errno"
	"github.com/A-Simple-Tictok-Project/douyin/repository"
)

type RelationshipCreateOrDeletePostFlow struct {
	// 请求需要的数据
	UserId      int64 `json:"user_id"`      // 女神ID
	SubscribeId int64 `json:"subscribe_id"` // 舔狗ID
	ActionType  int64 `json:"action_type"`
	// 响应需要的数据
}

func NewRelationshipCreateOrDeletePostFlow(userId, subscribeId, actionType int64) *RelationshipCreateOrDeletePostFlow {
	return &RelationshipCreateOrDeletePostFlow{
		UserId:      userId,
		SubscribeId: subscribeId,
		ActionType:  actionType,
	}
}
func RelationshipCreateOrDeletePost(userId, subscribeId, actionType int64) errno.ErrNo {
	return NewRelationshipCreateOrDeletePostFlow(userId, subscribeId, actionType).Do()
}

func (f *RelationshipCreateOrDeletePostFlow) Do() errno.ErrNo {
	// 1. check param
	if err := f.checkParam(); err.ErrCode != 0 {
		return err
	}
	// 2. create or delete relationship
	if err := f.createOrDeleteRelationship(); err.ErrCode != 0 {
		return err
	}
	return errno.Success
}

func (f *RelationshipCreateOrDeletePostFlow) checkParam() errno.ErrNo {
	if f.UserId == 0 || f.SubscribeId == 0 {
		return errno.ParamErr
	}
	switch f.ActionType {
	case constants.ActionDo: // 新增
	case constants.ActionNotDo: // 删除
	default:
		return errno.ParamErr
	}
	return errno.Success
}

func (f *RelationshipCreateOrDeletePostFlow) createOrDeleteRelationship() errno.ErrNo {
	relationship := &repository.Relationship{
		UserId:      f.UserId,
		SubscribeId: f.SubscribeId,
	}
	var err error
	switch f.ActionType {
	case constants.ActionDo: // 新增
		_, err = repository.NewRelationshipDaoInstance().CreateRelationship(relationship)
	case constants.ActionNotDo: // 删除
		err = repository.NewRelationshipDaoInstance().DeleteRelationship(relationship)
	default:
		return errno.ParamErr
	}
	if err != nil {
		return errno.ConvertErr(err)
	}
	return errno.Success
}
