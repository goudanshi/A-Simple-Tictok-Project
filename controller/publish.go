package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
