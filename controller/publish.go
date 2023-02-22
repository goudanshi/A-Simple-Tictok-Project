package controller

import (
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PublishListResponse struct {
	BaseResponse
	VideoList []service.VideoUser `json:"video_list"`
	NextTime  int64               `json:"next_time"`
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
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		buildError(c, err.Error())
		return
	}
	data, err := service.QueryUserVideo(userId)
	fmt.Println("publishlist")
	fmt.Println(data)
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
		return
	}
	c.JSON(http.StatusOK, PublishListResponse{
		BaseResponse: BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: data,
		NextTime:  data[len(data)-1].CreateDate.UnixNano(),
	})
}
