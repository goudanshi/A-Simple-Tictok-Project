package router

import (
	"github.com/A-Simple-Tictok-Project/douyin/handler"
	"github.com/A-Simple-Tictok-Project/douyin/middleware"
	"github.com/gin-gonic/gin"
)

var registerFavoriteRoute = func(r *gin.RouterGroup) {
	favoriteGroup := r.Group("/favorite")
	{
		favoriteGroup.POST("/action", middleware.JWTToken(), handler.LikeCreateOrDelete)
		favoriteGroup.GET("/list", handler.LikeQueryList)
	}
}
