package handler

import (
	"github.com/A-Simple-Tictok-Project/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PageLikeQueryListData struct {
	StatusCode int64                    `json:"status_code"`
	StatusMsg  string                   `json:"status_msg"`
	VideoList  []map[string]interface{} `json:"video_list,omitempty"`
}

func likeQueryListGet(userIdStr string) *PageLikeQueryListData {
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoInfo, err := service.LikeQueryList(userId)
	if err.ErrCode != 0 || videoInfo == nil {
		return &PageLikeQueryListData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}
	}
	videoList := make([]map[string]interface{}, 0)
	for _, video := range videoInfo.Videos {
		videoList = append(videoList, map[string]interface{}{
			"id": video.ID,
			"author": map[string]interface{}{
				"id":             videoInfo.AuthorMap[video.UserId].ID,
				"name":           videoInfo.AuthorMap[video.UserId].Username,
				"follow_count":   videoInfo.AuthorMap[video.UserId].UserCount, // 女神
				"follower_count": videoInfo.AuthorMap[video.UserId].SubscribeCount,      // 舔狗
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
	return &PageLikeQueryListData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
		VideoList:  videoList,
	}
}

func LikeQueryList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	data := likeQueryListGet(userIdStr)
	c.JSON(http.StatusOK, data)
}
