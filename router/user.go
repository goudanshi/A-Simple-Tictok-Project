package router

import (
	"github.com/gin-gonic/gin"
	"github.com/A-Simple-Tictok-Project/douyin/handler"
	"github.com/A-Simple-Tictok-Project/douyin/middleware"
)

var registerUserRoute = func(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
		userGroup.GET("/", middleware.JWTToken(), handler.Userinfo)
	}
}
