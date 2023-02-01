package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/A-Simple-Tictok-Project/douyin/service"
	"net/http"
	"strconv"
)

type PageLikeCreateOrDeleteData struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func likeCreateOrDeletePost(userIdStr, videoIdStr, actionTypeStr string) *PageLikeCreateOrDeleteData {
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	actionType, _ := strconv.ParseInt(actionTypeStr, 10, 64)
	err := service.LikeCreateOrDeletePost(userId, videoId, actionType)
	return &PageLikeCreateOrDeleteData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}

func LikeCreateOrDelete(c *gin.Context) {
	tempUserId, _ := c.Get("userId")
	userIdStr := tempUserId.(string)
	videoIdStr := c.PostForm("video_id")
	actionTypeStr := c.PostForm("action_type")
	data := likeCreateOrDeletePost(userIdStr, videoIdStr, actionTypeStr)
	c.JSON(http.StatusOK, data)
}
