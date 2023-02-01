package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/A-Simple-Tictok-Project/douyin/service"
	"net/http"
	"strconv"
)

type PageVideoCreateData struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func videoCreatePost(title, userIdStr string) *PageVideoCreateData {
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	// TODO 1. 保存视频信息 ——> 1.保存视频 2.截取封面
	playUrl := "http://127.0.0.1:8080/video/video.mp4"
	coverUrl := "http://127.0.0.1:8080/video/cover.jpg"
	// 2. 获取service层结果
	_, err := service.VideoCreatePost(title, playUrl, coverUrl, userId)
	return &PageVideoCreateData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}

func VideoCreate(c *gin.Context) {
	// TODO 视频数据的获取
	tempUserId, _ := c.Get("userId")
	userIdStr := tempUserId.(string)
	title := c.PostForm("title")
	data := videoCreatePost(title, userIdStr)
	c.JSON(http.StatusOK, data)
}
