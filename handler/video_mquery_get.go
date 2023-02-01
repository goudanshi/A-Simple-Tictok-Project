package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jason/douyin/service"
	"net/http"
)

type PageVideoMQueryData struct {
	StatusCode int64                    `json:"status_code"`
	StatusMsg  string                   `json:"status_msg"`
	NextTime   string                   `json:"next_time,omitempty"`
	VideoList  []map[string]interface{} `json:"video_list,omitempty"`
}

func videoMQueryGet(latestTime string) *PageVideoMQueryData {
	// 获取service层数据
	videoInfo, err := service.VideoMQueryGet(latestTime)
	if err.ErrCode != 0 || videoInfo == nil {
		return &PageVideoMQueryData{
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
				"follow_count":   videoInfo.AuthorMap[video.UserId].UserCount,      // 女神
				"follower_count": videoInfo.AuthorMap[video.UserId].SubscribeCount, // 舔狗
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
	return &PageVideoMQueryData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
		NextTime:   videoInfo.Videos[len(videoInfo.Videos)-1].CreatedAt.Format("2006-01-02 15:04:05"),
		VideoList:  videoList,
	}
}

func VideoMQuery(c *gin.Context) {
	latestTime := c.Query("latest_time")
	data := videoMQueryGet(latestTime)
	c.JSON(http.StatusOK, data)
}
