package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetFeedList(c *gin.Context) {

	data, err := service.VideoList(1)
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{1, err.Error()})
		return
	}
	c.JSON(http.StatusOK, struct {
		BaseResponse
		Video    []service.VideoUser `json:"video_list"`
		NextTime *time.Time          `json:"next_time"`
	}{BaseResponse{0, "success"}, data, &data[len(data)-1].CreateDate})
}

func GetFeed(c *gin.Context) {
	path := c.Query("name")
	reader, length, err := service.GetVideoFeed(path)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	c.DataFromReader(http.StatusOK, length, "video/mp4", reader, make(map[string]string))
}

func GetCover(c *gin.Context) {
	path := c.Query("name")
	reader, length, err := service.GetVideoFeed(path)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	c.DataFromReader(http.StatusOK, length, "image/png", reader, make(map[string]string))
}
