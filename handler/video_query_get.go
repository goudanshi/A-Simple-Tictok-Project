package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/A-Simple-Tictok-Project/douyin/service"
	"net/http"
	"strconv"
)

type PageVideoQueryData struct {
	StatusCode int64                    `json:"status_code"`
	StatusMsg  string                   `json:"status_msg"`
	VideoList  []map[string]interface{} `json:"video_list,omitempty"`
}

func videoQueryGet(userIdStr string) *PageVideoQueryData {
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	// 获取service层数据
	videoInfo, err := service.VideoQueryGet(userId)
	if err.ErrCode != 0 {
		return &PageVideoQueryData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}
	}
	videos := make([]map[string]interface{}, 0)
	for _, video := range videoInfo.Videos {
		videos = append(videos, map[string]interface{}{
			"id": video.ID,
			"author": map[string]interface{}{
				"id":             videoInfo.Author.ID,
				"name":           videoInfo.Author.Username,
				"follow_count":   videoInfo.Author.UserCount,      // 女神
				"follower_count": videoInfo.Author.SubscribeCount, // 舔狗
				"is_follow":      false,
			},
			"play_url":       video.PlayURL,
			"cover_url":      video.CoverURL,
			"favorite_count": video.LikeCount,
			"comment_count":  video.CommentCount,
			"is_favorite":    false,
			"title":          video.Title,
		})
	}
	return &PageVideoQueryData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
		VideoList:  videos,
	}
}

func VideoQuery(c *gin.Context) {
	userIdStr := c.Query("user_id")
	data := videoQueryGet(userIdStr)
	c.JSON(http.StatusOK, data)
}
