package handler

import (
	"github.com/A-Simple-Tictok-Project/douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PageRelationshipCreateOrDeleteData struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func relationshipCreateOrDeletePost(userIdStr, subscribeIdStr, actionTypeStr string) *PageRelationshipCreateOrDeleteData {
	// userIdStr:女神ID
	// subscribeIdStr:舔狗ID
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	subscribeId, _ := strconv.ParseInt(subscribeIdStr, 10, 64)
	actionType, _ := strconv.ParseInt(actionTypeStr, 10, 64)
	// 获取service层数据
	err := service.RelationshipCreateOrDeletePost(userId, subscribeId, actionType)
	return &PageRelationshipCreateOrDeleteData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}

func RelationshipCreateOrDelete(c *gin.Context) {
	tempSubscribeId, _ := c.Get("userId")
	subscribeIdStr := tempSubscribeId.(string) // 当前登录的用户ID，也是舔狗
	userIdStr := c.PostForm("to_user_id")      // 女神ID
	actionTypeStr := c.PostForm("action_type")
	data := relationshipCreateOrDeletePost(userIdStr, subscribeIdStr, actionTypeStr)
	c.JSON(http.StatusOK, data)
}
