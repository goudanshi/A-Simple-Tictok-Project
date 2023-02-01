package router

import (
	"github.com/gin-gonic/gin"
	"github.com/A-Simple-Tictok-Project/douyin/handler"
	"github.com/A-Simple-Tictok-Project/douyin/middleware"
)

var registerFavoriteRoute = func(r *gin.RouterGroup) {
	favoriteGroup := r.Group("/favorite")
	{
		favoriteGroup.POST("/action", middleware.JWTToken(), handler.LikeCreateOrDelete)
		favoriteGroup.GET("/list")
	}
}
