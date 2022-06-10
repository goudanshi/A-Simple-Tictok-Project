package controller

import (
	"douyin/repository"
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type relationActionRequest struct {
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type"`
}

func RelationAction(c *gin.Context) {
	var request relationActionRequest
	c.ShouldBind(&request)
	userId := service.GetUserId(c)
	switch request.ActionType {
	case util.NEW_RELATION:
		_, err := service.NewRelation(&repository.Relation{
			FollowId:   request.ToUserId,
			FollowerId: userId,
		})
		if err != nil {
			buildError(c, err.Error())
			return
		}
		buildSuccess(c)
		return
	case util.DELETE_RELATION:
		err := service.DeleteRelation(request.ToUserId, userId)
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
	userId := service.GetUserId(c)
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
	userId := service.GetUserId(c)
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
