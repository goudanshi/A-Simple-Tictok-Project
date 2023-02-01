package router

import (
	"github.com/A-Simple-Tictok-Project/douyin/handler"
	"github.com/A-Simple-Tictok-Project/douyin/middleware"
	"github.com/gin-gonic/gin"
)

var registerRelationRoute = func(r *gin.RouterGroup) {
	relationGroup := r.Group("/relation")
	{
		relationGroup.POST("/action", middleware.JWTToken(), handler.RelationshipCreateOrDelete)
		relationGroup.GET("/follow/list", handler.RelationshipQueryNList)
		relationGroup.GET("/follower/list", handler.RelationshipQueryTList)
	}
}
