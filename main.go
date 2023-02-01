package main

import (
	"github.com/jason/douyin/repository"
	"github.com/jason/douyin/router"
	"github.com/jason/douyin/utils"
	"os"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	engine := router.Init()
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	if err := utils.InitLogger(); err != nil {
		return err
	}
	return nil
}
