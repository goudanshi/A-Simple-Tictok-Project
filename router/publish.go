package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jason/douyin/handler"
	"github.com/jason/douyin/middleware"
)

var registerVideoRoute = func(r *gin.RouterGroup) {
	videoGroup := r.Group("/publish")
	{
		videoGroup.POST("/action", middleware.JWTToken(), handler.VideoCreate)
		videoGroup.GET("/list", handler.VideoQuery)
	}
	r.GET("/feed", handler.VideoMQuery)
}
