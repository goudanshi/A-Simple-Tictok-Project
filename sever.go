package main

import (
	"douyin/repository"
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init(); err != nil {
		util.Logger.Error(err.Error())
		os.Exit(1)
	}

	r := gin.Default()
	initRouter(r)
	err := r.Run(":8100")
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	if err := service.AuthMiddlewareInit(); err != nil {
		return err
	}
	//if err := util.InitRedis(); err != nil {
	//	return err
	//}
	if err := util.InitOSS(); err != nil {
		return err
	}
	return nil
}

//
