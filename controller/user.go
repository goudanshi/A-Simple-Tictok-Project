package controller

import (
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type registerResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
	BaseResponse
}

type userInfoResponse struct {
	BaseResponse
	User userResponse `json:"user"`
}

type userResponse struct {
	Id            int64  `json:"id"`
	Username      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func Login(c *gin.Context) {
	service.GetAuthInstance().LoginHandler(c)
}

func Register(c *gin.Context) {
	var request repository.UserRequest
	c.ShouldBind(&request)
	id, token, err := service.Register(&repository.User{
		Username: request.Username,
		Password: request.Password,
	})

	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, registerResponse{
		UserId: id,
		Token:  token,
		BaseResponse: BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
	})
}

func UserInfo(c *gin.Context) {
	tmp := service.GetUserId(c)
	idRaw := c.Query("user_id")
	id, _ := strconv.ParseInt(idRaw, 10, 64)
	user, err := service.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	isFollow := service.IsFollow(id, tmp)

	c.JSON(http.StatusOK, userInfoResponse{
		BaseResponse: BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		User: userResponse{
			Id:            user.Id,
			Username:      user.Username,
			FollowerCount: user.FollowerCount,
			FollowCount:   user.FollowCount,
			IsFollow:      isFollow,
		},
	})
}

func Auth(c *gin.Context) {
	service.Auth(c)
}
