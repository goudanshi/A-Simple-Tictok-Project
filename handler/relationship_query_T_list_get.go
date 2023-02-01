package handler

import (
	"github.com/jason/douyin/service"
	"strconv"
)

type PageRelationshipQueryTListData struct {
	StatusCode int64                    `json:"status_code"`
	StatusMsg  string                   `json:"status_msg"`
	UserList   []map[string]interface{} `json:"user_list,omitempty"`
}

func RelationshipQueryTList(userIdStr string) *PageRelationshipQueryTListData {
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	// 获取service层数据
	relationshipInfo, err := service.RelationshipQueryTListGet(userId)
	if err.ErrCode != 0 {
		return &PageRelationshipQueryTListData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}
	}
	userList := make([]map[string]interface{}, 0)
	for _, subscribe := range relationshipInfo.SubscribeUsers {
		userList = append(userList, map[string]interface{}{
			"id":             subscribe.ID,
			"name":           subscribe.Username,
			"follow_count":   subscribe.UserCount,      // 女神
			"follower_count": subscribe.SubscribeCount, // 舔狗
			"is_follow":      false,
		})
	}
	return &PageRelationshipQueryTListData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
		UserList:   userList,
	}
}
