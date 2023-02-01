package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())
	v1 := r.Group("/douyin")
	{
		// 注册用户相关路由
		registerUserRoute(v1)
		// 注册视频相关路由
		registerVideoRoute(v1)
		// 注册点赞相关路由
		registerFavoriteRoute(v1)
		// 注册评论模块相关路由
		registerCommentRoute(v1)
		// 注册社交模块相关路由
		registerRelationRoute(v1)
	}

	return r
}
