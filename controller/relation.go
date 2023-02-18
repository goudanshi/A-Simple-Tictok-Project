package controller

import (
	"douyin/repository"
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type relationActionRequest struct {
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type"`
}

func RelationAction(c *gin.Context) {

	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	userId := service.GetUserId(c)
	switch actionType {
	case util.NEW_RELATION:
		_, err := service.NewRelation(&repository.Relation{
			FollowId:   toUserId,
			FollowerId: userId,
		})
		if err != nil {
			buildError(c, err.Error())
			return
		}
		buildSuccess(c)
		return
	case util.DELETE_RELATION:
		err := service.DeleteRelation(toUserId, userId)
		if err != nil {
			buildError(c, err.Error())
			return
		}
		buildSuccess(c)
		return
	default:
		buildError(c, "error action type")
		return
	}
}

func FollowList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	data, err := service.GetFollowList(userId)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, struct {
		BaseResponse
		UserList []service.User `json:"user_list"`
	}{
		BaseResponse: BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: data,
	})
}

func FollowerList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	data, err := service.GetFollowerList(userId)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, struct {
		BaseResponse
		UserList []service.User `json:"user_list"`
	}{
		BaseResponse: BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: data,
	})
}
