package handler

import (
	"github.com/jason/douyin/constants"
	"github.com/jason/douyin/service"
	"strconv"
)

type PageCommentCreateOrDeleteData struct {
	StatusCode int64                  `json:"status_code"`
	StatusMsg  string                 `json:"status_msg"`
	Comment    map[string]interface{} `json:"comment,omitempty"`
}

func CommentCreateOrDelete(userIdStr, videoIdStr, commentIdStr, actionTypeStr, content string) *PageCommentCreateOrDeleteData {
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)
	actionType, _ := strconv.ParseInt(actionTypeStr, 10, 64)
	// 获取service层数据
	commentInfo, err := service.CommentCreateOrDeletePost(userId, videoId, commentId, actionType, content)
	if err.ErrCode != 0 {
		return &PageCommentCreateOrDeleteData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}
	}
	if actionType == constants.ActionDo {
		return &PageCommentCreateOrDeleteData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
			Comment: map[string]interface{}{
				"id": commentInfo.CommentId,
				"user": map[string]interface{}{
					"id":             commentInfo.User.ID,
					"name":           commentInfo.User.Username,
					"follow_count":   commentInfo.User.UserCount,      // 女神
					"follower_count": commentInfo.User.SubscribeCount, // 舔狗
					"is_follow":      true,
				},
				"content":     commentInfo.Comment.Content,
				"create_date": commentInfo.Comment.CreatedAt,
			},
		}
	}
	return &PageCommentCreateOrDeleteData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}
