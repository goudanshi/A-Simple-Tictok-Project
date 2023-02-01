package main

import (
	"github.com/A-Simple-Tictok-Project/douyin/repository"
	"github.com/A-Simple-Tictok-Project/douyin/router"
	"github.com/A-Simple-Tictok-Project/douyin/utils"
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
