package controller

import (
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"net/http"
)

func GetFeedList(c *gin.Context) {

	var userId int64
	userId = -1
	tokenString := c.Query("token")
	if tokenString != "" {
		token, _ := service.GetAuthInstance().ParseTokenString(tokenString)
		for key, value := range token.Claims.(jwt.MapClaims) {
			if key == "identity" {
				userId = int64(int(value.(float64)))
			}
		}
	}

	data, err := service.VideoList(userId)
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{1, err.Error()})
		return
	}
	c.JSON(http.StatusOK, struct {
		BaseResponse
		Video    []service.VideoUser `json:"video_list"`
		NextTime int64               `json:"next_time"`
	}{BaseResponse{0, "success"}, data, data[len(data)-1].CreateDate.UnixNano()})
}

func GetFeed(c *gin.Context) {
	path := c.Query("name")
	reader, length, _, err := service.GetVideoFeed(path)
	if err != nil {
		buildError(c, err.Error())
		return
	}

	//bytes, _ := ioutil.ReadAll(reader)
	fmt.Println(length)
	//c.Data(http.StatusOK, "video/mp4", bytes)
	c.Header("Accept-Ranges", "bytes")
	//c.Header("Last-Modified", "Sun, 12 Jun 2022 15:38:23 UTC")
	c.DataFromReader(http.StatusOK, length, "video/mp4", reader, nil)
	c.Next()

}

func GetCover(c *gin.Context) {
	path := c.Query("name")
	reader, length, _, err := service.GetVideoCover(path)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	fmt.Println(length)
	bytes, _ := ioutil.ReadAll(reader)
	c.Data(http.StatusOK, "image/png", bytes)
	//c.DataFromReader(http.StatusOK, length, "image/png", reader, make(map[string]string))

}
