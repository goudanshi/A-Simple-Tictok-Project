package controller

import (
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type favoriteActionRequest struct {
	VideoId    int64 `form:video_id`
	ActionType int32 `form:action_type`
}

func FavoriteAction(c *gin.Context) {
	var request favoriteActionRequest
	c.ShouldBind(&request)
	userId := service.GetUserId(c)
	if request.ActionType == util.EXEC_FAVORITE {
		_, err := service.NewFavorite(userId, request.VideoId)
		if err != nil {
			buildError(c, err.Error())
			return
		}
		buildSuccess(c)
		return
	} else if request.ActionType == util.CANCEL_FAVORITE {
		err := service.CancelFavorite(userId, request.VideoId)
		if err != nil {
			buildError(c, err.Error())
			return
		}
		buildSuccess(c)
		return
	} else {
		buildError(c, "error action type")
		return
	}
}

func FavoriteList(c *gin.Context) {
	userId := service.GetUserId(c)

	data, err := service.QueryFavoriteVideo(userId)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, struct {
		BaseResponse
		VideoList []service.VideoUser `json:"video_list"`
	}{
		BaseResponse: BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: data,
	})
}
