package router

import (
	"github.com/A-Simple-Tictok-Project/douyin/handler"
	"github.com/A-Simple-Tictok-Project/douyin/middleware"
	"github.com/gin-gonic/gin"
)

var registerCommentRoute = func(r *gin.RouterGroup) {
	commentGroup := r.Group("/comment")
	{
		commentGroup.POST("/action", middleware.JWTToken(), handler.CommentCreateOrDelete)
		commentGroup.GET("/list", handler.CommentQuery)
	}
}
