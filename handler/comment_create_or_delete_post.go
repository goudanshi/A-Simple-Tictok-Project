package handler

import (
	"github.com/A-Simple-Tictok-Project/douyin/constants"
	"github.com/A-Simple-Tictok-Project/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PageCommentCreateOrDeleteData struct {
	StatusCode int64                  `json:"status_code"`
	StatusMsg  string                 `json:"status_msg"`
	Comment    map[string]interface{} `json:"comment,omitempty"`
}

func commentCreateOrDelete(userIdStr, videoIdStr, commentIdStr, actionTypeStr, content string) *PageCommentCreateOrDeleteData {
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
				"id": commentInfo.Comment.ID,
				"user": map[string]interface{}{
					"id":             commentInfo.User.ID,
					"name":           commentInfo.User.Username,
					"follow_count":   commentInfo.User.UserCount,      // 女神
					"follower_count": commentInfo.User.SubscribeCount, // 舔狗
					"is_follow":      false,
				},
				"content":     commentInfo.Comment.Content,
				"create_date": commentInfo.Comment.CreatedAt.Format("2006-01-02 15:04:05"),
			},
		}
	}
	return &PageCommentCreateOrDeleteData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}

func CommentCreateOrDelete(c *gin.Context) {
	tempUserId, _ := c.Get("userId")
	userIdStr := tempUserId.(string)
	videoIdStr := c.PostForm("video_id")
	actionTypeStr := c.PostForm("action_type")
	commentIdStr := c.PostForm("comment_id")
	commentTextStr := c.PostForm("comment_text")
	data := commentCreateOrDelete(userIdStr, videoIdStr, commentIdStr, actionTypeStr, commentTextStr)
	c.JSON(http.StatusOK, data)
}
