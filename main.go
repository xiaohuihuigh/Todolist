package main

import (
	"todoList/app/application"
	"todoList/app/models"
	"todoList/app/route"
	"github.com/gin-gonic/gin"
)

func main(){
	models.Setup()
	r := application.ApplicationInif.Init()
	route.SetRouter(r.Engine)

	en := gin.Default()
	route.SetRouter(en)
	//en.Run()
	r.Run() // 在 0.0.0.0:8080 上监听并服务
	//web.Run("Minimal webview example","https://www.baidu.com")
}
//func main() {
//	debug := true
//	w := webview.New(debug)
//	defer w.Destroy()
//	w.SetTitle("Minimal webview example")
//	w.SetSize(800, 600, webview.HintNone)
//	w.Navigate("https://www.baidu.com")
//	w.Run()
//}