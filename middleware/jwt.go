package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/A-Simple-Tictok-Project/douyin/constants"
	. "github.com/A-Simple-Tictok-Project/douyin/utils"
	"net/http"
	"strconv"
)

func JWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 0. 判断请求方法，因为token存放的位置不同：GET在请求地址中；POST在请求体中
		// 1. 获取token信息。
		var tokenString string
		if c.Request.Method == "GET" {
			tokenString = c.Query("token")
		}
		if c.Request.Method == "POST" {
			tokenString = c.PostForm("token")
		}
		// 2. 验证token信息
		jwtToken := NewJWT([]byte(constants.JWTPrivateKey))
		claims, err := jwtToken.ParseToken(tokenString)
		// 3. 验证失败，直接打回
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 10006,
				"status_msg":  err.Error(),
			})
			c.Abort() // 直接中断请求打回去
			return
		}
		fmt.Println(claims)
		// 4. 通过验证，给Context设置数据
		c.Set("userId", strconv.Itoa(int(claims.UserId))) // 太粗糙啦！！！
		c.Next()                                          // 往下走
	}
}
