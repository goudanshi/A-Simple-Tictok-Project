package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PublishListResponse struct {
	BaseResponse
	VideoList []service.VideoUser `json:"video_list"`
	NextTime  time.Time           `json:"next_time"`
}

func Publish(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetUserId(c))
	file, _ := c.FormFile("data")
	title := c.PostForm("title")
	_, err := service.PublishVideo(file, title, service.GetUserId(c))
	if err != nil {
		buildError(c, err.Error())
		return
	}
	buildSuccess(c)
}

func PublishList(c *gin.Context) {
	userId := service.GetUserId(c)
	data, err := service.QueryUserVideo(userId)
	if err != nil {
		buildError(c, err.Error())
	}
	if len(data) == 0 {
		c.JSON(http.StatusOK, PublishListResponse{
			BaseResponse: BaseResponse{
				StatusCode: 0,
				StatusMsg:  "success",
			},
		})
	}
	c.JSON(http.StatusOK, PublishListResponse{
		BaseResponse: BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: data,
		NextTime:  data[len(data)-1].CreateDate,
	})
}
