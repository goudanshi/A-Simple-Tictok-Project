package handler

import (
	"github.com/jason/douyin/service"
	"strconv"
)

type PageRelationshipCreateOrDeleteData struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func RelationshipCreateOrDelete(userIdStr, subscribeIdStr, actionTypeStr string) *PageRelationshipCreateOrDeleteData {
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
