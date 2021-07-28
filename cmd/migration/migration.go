package main

import "github.com/siconghe/blog/models"

func main() {
	//创建Tag表
	models.InitDB()
	(&models.Tag{}).CreateTable()
	(&models.Article{}).CreateTable()
}

