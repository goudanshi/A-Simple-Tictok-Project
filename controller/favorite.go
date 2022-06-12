package controller

import (
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type favoriteActionRequest struct {
	VideoId    int64 `form:video_id`
	ActionType int32 `form:action_type`
}

func FavoriteAction(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	userId := service.GetUserId(c)
	if actionType == util.EXEC_FAVORITE {
		_, err := service.NewFavorite(userId, videoId)
		if err != nil {
			buildError(c, err.Error())
			return
		}
		buildSuccess(c)
		return
	} else if actionType == util.CANCEL_FAVORITE {
		err := service.CancelFavorite(userId, videoId)
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
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		buildError(c, err.Error())
		return
	}

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
