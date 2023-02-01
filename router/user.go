package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jason/douyin/handler"
	"github.com/jason/douyin/middleware"
)

var registerUserRoute = func(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
		userGroup.GET("/", middleware.JWTToken(), handler.Userinfo)
	}
}
