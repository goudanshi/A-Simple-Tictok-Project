package main

import (
	"douyin/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	douyinGroup := r.Group("/douyin")
	{
		douyinGroup.GET("/feed/", controller.GetFeedList)
		douyinGroup.GET("/feed/get", controller.GetFeed)
		douyinGroup.GET("/feed/get/cover", controller.GetCover)

		userGroup := douyinGroup.Group("/user")
		{
			userGroup.POST("/register/", controller.Register)
			userGroup.POST("/login/", controller.Login)
			userGroup.GET("/", controller.Auth, controller.UserInfo)
		}

		publishGroup := douyinGroup.Group("/publish")
		{
			publishGroup.POST("/action/", controller.Publish)
			publishGroup.GET("/list/", controller.Auth, controller.PublishList)
		}

		favoriteGroup := douyinGroup.Group("/favorite", controller.Auth)
		{
			favoriteGroup.POST("/action/", controller.FavoriteAction)
			favoriteGroup.GET("/list/", controller.FavoriteList)
		}

		commentGroup := douyinGroup.Group("/comment", controller.Auth)
		{
			commentGroup.POST("/action/", controller.CommentAction)
			commentGroup.GET("/list/", controller.CommentList)
		}

		relationGroup := douyinGroup.Group("/relation", controller.Auth)
		{
			relationGroup.POST("/action/", controller.RelationAction)
			relationGroup.GET("/follow/list/", controller.FollowList)
			relationGroup.GET("/follower/list/", controller.FollowerList)
		}
	}
}
