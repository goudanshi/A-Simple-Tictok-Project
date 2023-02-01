package handler

import (
	"github.com/A-Simple-Tictok-Project/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PageUserQueryData struct {
	StatusCode int64                  `json:"status_code"`
	StatusMsg  string                 `json:"status_msg"`
	User       map[string]interface{} `json:"user,omitempty"`
}

func userQueryGet(userIdStr string) *PageUserQueryData {
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	// 获取service层结果
	users, err := service.UserQueryGet(userId)
	if err.ErrCode != 0 {
		return &PageUserQueryData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}
	}
	if len(users) == 0 {
		return &PageUserQueryData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}
	}
	return &PageUserQueryData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
		User: map[string]interface{}{
			"id":             users[0].ID,
			"name":           users[0].Username,
			"follow_count":   users[0].UserCount,      // 女神
			"follower_count": users[0].SubscribeCount, // 舔狗
			"is_follow":      false,
		},
	}
}

func Userinfo(c *gin.Context) {
	userIdStr := c.Query("user_id")
	data := userQueryGet(userIdStr)
	c.JSON(http.StatusOK, data)
}
