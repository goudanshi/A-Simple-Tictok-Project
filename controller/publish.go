package controller

import (
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type PublishListResponse struct {
	BaseResponse
	VideoList []service.VideoUser `json:"video_list"`
	NextTime  time.Time           `json:"next_time"`
}

func Publish(c *gin.Context) {

	tokenString := c.PostForm("token")
	token, _ := service.GetAuthInstance().ParseTokenString(tokenString)
	var userId int64
	for key, value := range token.Claims.(jwt.MapClaims) {
		if key == "identity" {
			userId = int64(int(value.(float64)))
		}
	}
	fmt.Println(userId)

	//c.JSON(http.StatusOK, service.GetUserId(c))
	file, _ := c.FormFile("data")
	title := c.PostForm("title")
	_, err := service.PublishVideo(file, title, userId)
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
