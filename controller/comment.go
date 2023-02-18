package controller

import (
	"douyin/repository"
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type commonActionRequest struct {
	VideoId     int64  `form:"video_id"`
	ActionType  int32  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

func CommentAction(c *gin.Context) {
	var request commonActionRequest
	c.ShouldBind(&request)
	userId := service.GetUserId(c)
	switch request.ActionType {
	case util.PUBLISH_COMMENT:
		_, err := service.NewComment(&repository.Comment{
			UserId:     userId,
			VideoId:    request.VideoId,
			Content:    request.CommentText,
			CreateDate: time.Now(),
		})
		if err != nil {
			buildError(c, err.Error())
			return
		}
		buildSuccess(c)
		return
	case util.DELETE_COMMENT:
		err := service.DeleteComment(request.CommentId)
		if err != nil {
			buildError(c, err.Error())
			return
		}
		buildSuccess(c)
		return
	default:
		buildError(c, "error action type")
	}
}

func CommentList(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	userId := service.GetUserId(c)

	data, err := service.GetCommentListByVideo(videoId, userId)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, struct {
		BaseResponse
		CommentList []service.Comment `json:"comment_list"`
	}{
		BaseResponse: BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		CommentList: data,
	})
}
