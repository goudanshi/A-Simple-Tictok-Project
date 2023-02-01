package service

import (
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/repository"
)

type CommentQueryGetFlow struct {
	// 请求需要的数据
	VideoId int64 `json:"video_id"`
	// 响应需要的数据
	Comments []*repository.Comment      `json:"comments"`
	UserMap  map[int64]*repository.User `json:"user_map"`
}

func NewCommentQueryGetFlow(videoId int64) *CommentQueryGetFlow {
	return &CommentQueryGetFlow{VideoId: videoId}
}

func CommentQueryGet(videoId int64) (*CommentQueryGetFlow, errno.ErrNo) {
	return NewCommentQueryGetFlow(videoId).Do()
}

func (f *CommentQueryGetFlow) Do() (*CommentQueryGetFlow, errno.ErrNo) {
	// 1. check param
	if err := f.checkParam(); err.ErrCode != 0 {
		return nil, err
	}
	// 2. query comment list information
	if err := f.queryComment(); err.ErrCode != 0 {
		return nil, err
	}
	// 3. query comment author list information
	if err := f.queryUser(); err.ErrCode != 0 {
		return nil, err
	}
	return f, errno.Success
}

func (f *CommentQueryGetFlow) checkParam() errno.ErrNo {
	if f.VideoId == 0 {
		return errno.ParamErr
	}
	return errno.Success
}
func (f *CommentQueryGetFlow) queryComment() errno.ErrNo {
	comments, err := repository.NewCommentDaoInstance().QueryCommentByVideoId(f.VideoId)
	if err != nil {
		return errno.ConvertErr(err)
	}
	f.Comments = comments
	return errno.Success
}

func (f *CommentQueryGetFlow) queryUser() errno.ErrNo {
	userIds := make([]int64, 0)
	for _, comment := range f.Comments {
		userIds = append(userIds, comment.UserId)
	}
	users, err := repository.NewUserDaoInstance().MQueryUserById(userIds)
	if err != nil {
		return errno.ConvertErr(err)
	}
	userMap := make(map[int64]*repository.User)
	for _, user := range users {
		userMap[user.ID] = user
	}
	f.UserMap = userMap
	return errno.Success
}
