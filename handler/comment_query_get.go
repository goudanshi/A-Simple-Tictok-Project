package handler

import (
	"github.com/A-Simple-Tictok-Project/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PageCommentQueryGetData struct {
	StatusCode  int64                    `json:"status_code"`
	StatusMsg   string                   `json:"status_msg"`
	CommentList []map[string]interface{} `json:"comment_list"`
	//CommentList []map[string]interface{} `json:"comment_list,omitempty"`
}

func commentQueryGet(videoIdStr string) *PageCommentQueryGetData {
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	// 获取service层的数据
	commentInfo, err := service.CommentQueryGet(videoId)
	if err.ErrCode != 0 {
		return &PageCommentQueryGetData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}
	}
	var commentList []map[string]interface{}
	for _, comment := range commentInfo.Comments {
		commentList = append(commentList, map[string]interface{}{
			"id": comment.ID,
			"user": map[string]interface{}{
				"id":             commentInfo.UserMap[comment.UserId].ID,
				"name":           commentInfo.UserMap[comment.UserId].Username,
				"follow_count":   commentInfo.UserMap[comment.UserId].UserCount,      // 女神
				"follower_count": commentInfo.UserMap[comment.UserId].SubscribeCount, // 舔狗
				"is_follow":      false,
			},
			"content":     comment.Content,
			"create_date": comment.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &PageCommentQueryGetData{
		StatusCode:  err.ErrCode,
		StatusMsg:   err.ErrMsg,
		CommentList: commentList,
	}
}

func CommentQuery(c *gin.Context) {
	videoIdStr := c.Query("video_id")
	data := commentQueryGet(videoIdStr)
	c.JSON(http.StatusOK, data)
}
