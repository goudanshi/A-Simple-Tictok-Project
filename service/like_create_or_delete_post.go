package service

import (
	"github.com/jason/douyin/constants"
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/repository"
)

type LikeCreateOrDeletePostFlow struct {
	// 请求需要的数据
	UserId     int64 `json:"user_id"`
	VideoId    int64 `json:"video_id"`
	ActionType int64 `json:"action_type"`
	// 响应需要的数据
	//LikeId int64 `json:"like_id"`
}

func NewLikeCreateOrDeletePostFlow(userId, videoId, actionType int64) *LikeCreateOrDeletePostFlow {
	return &LikeCreateOrDeletePostFlow{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: actionType,
	}
}
func LikeCreateOrDeletePost(userId, videoId, actionType int64) errno.ErrNo {
	return NewLikeCreateOrDeletePostFlow(userId, videoId, actionType).Do()
}

func (f *LikeCreateOrDeletePostFlow) Do() errno.ErrNo {
	// 1. check param
	if err := f.checkParam(); err.ErrCode != 0 {
		return err
	}
	// 2. create or delete like information
	if err := f.createOrDeleteLike(); err.ErrCode != 0 {
		return err
	}
	return errno.Success
}
func (f *LikeCreateOrDeletePostFlow) checkParam() errno.ErrNo { // 无论是新增还是删除，都需要userId和videoId
	switch f.ActionType {
	case constants.ActionDo: // 新增
	case constants.ActionNotDo: // 删除
	default:
		return errno.ParamErr
	}
	if f.UserId == 0 || f.VideoId == 0 {
		return errno.ParamErr
	}
	return errno.Success
}

func (f *LikeCreateOrDeletePostFlow) createOrDeleteLike() errno.ErrNo {
	like := &repository.Like{
		UserId:  f.UserId,
		VideoId: f.VideoId,
	}
	var err error
	switch f.ActionType {
	case constants.ActionDo: // 新增
		err = repository.NewLikeDaoInstance().CreateLike(like)
	case constants.ActionNotDo: // 删除
		err = repository.NewLikeDaoInstance().DeleteLike(like)
	}
	if err != nil {
		return errno.ConvertErr(err)
	}
	return errno.Success
}
