package main

import "github.com/siconghe/blog/models"

func main() {
	//创建Tag表
	(&models.Tag{}).CreateTable()
}

