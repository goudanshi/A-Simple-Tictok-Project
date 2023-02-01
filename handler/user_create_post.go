package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jason/douyin/constants"
	"github.com/jason/douyin/errno"
	"github.com/jason/douyin/service"
	"github.com/jason/douyin/utils"
	"net/http"
	"time"
)

type PageUserCreateData struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}

func userCreatePost(username, password string) *PageUserCreateData {
	// 获取service层结果
	userId, err := service.UserCreatePost(username, password)
	if err.ErrCode != 0 {
		return &PageUserCreateData{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}
	}
	// 签发token信息
	jwtToken := utils.NewJWT([]byte(constants.JWTPrivateKey))
	claims := &utils.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Audience:  "字节跳动",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 7天的过期时间
			IssuedAt:  time.Now().Unix(),                         // token发布时间
			Issuer:    "抖音",
			//NotBefore: time.Now().Add(time.Hour).Unix(), // token开始生效时间，是以IssuedAt的时间开始计算
		},
	}
	token, err1 := jwtToken.CreateToken(claims) // 这里为什么要重新命名一个err1呢？因为我们前面还有一个err，两个err的类型不一样，所以不能共用
	if err1 != nil {
		utils.Logger.Error("generate token failed " + err1.Error())
		return &PageUserCreateData{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  errno.ServiceErr.ErrMsg,
			UserId:     userId,
			Token:      token,
		}
	}
	return &PageUserCreateData{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
		UserId:     userId,
		Token:      token,
	}
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	data := userCreatePost(username, password)
	c.JSON(http.StatusOK, data)
}