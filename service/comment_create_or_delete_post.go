package service

import (
	"github.com/jason/douyin/constants"
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/repository"
)

type CommentCreateOrDeletePostFlow struct {
	// 请求需要的数据
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	CommentId  int64  `json:"comment_id"`
	ActionType int64  `json:"action_type"`
	Content    string `json:"content"`
	// 响应需要的数据
	Comment *repository.Comment `json:"comment"`
	User    *repository.User    `json:"user"`
}

func NewCommentCreateOrDeletePostFlow(userId, videoId, commentId, actionType int64, content string) *CommentCreateOrDeletePostFlow {
	return &CommentCreateOrDeletePostFlow{
		UserId:     userId,
		VideoId:    videoId,
		CommentId:  commentId,
		ActionType: actionType,
		Content:    content,
	}
}
func CommentCreateOrDeletePost(userId, videoId, commentId, actionType int64, content string) (*CommentCreateOrDeletePostFlow, errno.ErrNo) {
	return NewCommentCreateOrDeletePostFlow(userId, videoId, commentId, actionType, content).Do()
}

func (f *CommentCreateOrDeletePostFlow) Do() (*CommentCreateOrDeletePostFlow, errno.ErrNo) {
	// 1. check param
	if err := f.checkParam(); err.ErrCode != 0 {
		return nil, err
	}
	// 2. do create or delete
	if err := f.createOrDeleteComment(); err.ErrCode != 0 {
		return nil, err
	}
	// 3. query comment author
	if err := f.queryUser(); err.ErrCode != 0 {
		return nil, err
	}
	return f, errno.Success
}
func (f *CommentCreateOrDeletePostFlow) checkParam() errno.ErrNo {
	switch f.ActionType {
	case constants.ActionDo: // 新增
		if f.UserId != 0 || f.VideoId != 0 || f.Content != "" {
			return errno.Success
		}
	case constants.ActionNotDo: // 删除
		if f.CommentId != 0 {
			return errno.Success
		}
	}
	return errno.ParamErr
}

func (f *CommentCreateOrDeletePostFlow) createOrDeleteComment() errno.ErrNo {
	comment := &repository.Comment{
		Content: f.Content,
		UserId:  f.UserId,
		VideoId: f.VideoId,
	}
	var err error
	switch f.ActionType {
	case constants.ActionDo: // 新增
		err = repository.NewCommentDaoInstance().CreateComment(comment)
	case constants.ActionNotDo: // 删除
		err = repository.NewCommentDaoInstance().DeleteComment(comment)
	}
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.Comment = comment
	return errno.Success
}

func (f *CommentCreateOrDeletePostFlow) queryUser() errno.ErrNo {
	users, err := repository.NewUserDaoInstance().MQueryUserById([]int64{f.Comment.UserId})
	if err != nil {
		return errno.ConvertErr(err)
	}
	if len(users) == 0 {
		return errno.UserNotExistErr
	}
	f.User = users[0]
	return errno.Success
}
