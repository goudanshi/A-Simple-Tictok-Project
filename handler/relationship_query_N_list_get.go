package handler

import (
	"github.com/A-Simple-Tictok-Project/douyin/service"
	"strconv"
)

type PageRelationshipQueryNListData struct {
	StatusCode int64                    `json:"status_code"`
	StatusMsg  string                   `json:"status_msg"`
	UserList   []map[string]interface{} `json:"user_list,omitempty"`
}

func RelationshipQueryNList(subscribeIdStr string) *PageRelationshipQueryNListData {
	subscribeId, _ := strconv.ParseInt(subscribeIdStr, 10, 64)
	// 获取service层数据
	relationshipInfo, err := service.RelationshipQueryNListGet(subscribeId)
	if err.ErrCode != 0 {
		return &PageRelationshipQueryNListData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}
	}
	userList := make([]map[string]interface{}, 0)
	for _, subscribe := range relationshipInfo.Users {
		userList = append(userList, map[string]interface{}{
			"id":             subscribe.ID,
			"name":           subscribe.Username,
			"follow_count":   subscribe.UserCount,      // 女神
			"follower_count": subscribe.SubscribeCount, // 舔狗
			"is_follow":      false,
		})
	}
	return &PageRelationshipQueryNListData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
		UserList:   userList,
	}
}
