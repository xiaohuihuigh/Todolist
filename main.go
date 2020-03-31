package main

import (
	"git.qutoutiao.net/todoList/app/application"
	"git.qutoutiao.net/todoList/app/route"
	"github.com/gin-gonic/gin"
)

func main(){
	r := application.ApplicationInif.Init()
	route.SetRouter(r.Engine)

	en := gin.Default()
	route.SetRouter(en)
	//en.Run()
	r.Run() // 在 0.0.0.0:8080 上监听并服务
}
