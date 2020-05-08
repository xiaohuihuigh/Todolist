package route

import (
	"todoList/app/components/utils"
	"todoList/app/controllers"
	"github.com/gin-gonic/gin"
)

//SetRouter 完全按gin的方式设置路由
func SetRouter(engine *gin.Engine) {
	//engine.Use(gin.Recovery(), gin.Logger())

	//engine.NoRoute(func(ctx *gin.Context) {
	//	utils.Response(ctx, utils.ErrorString{Code: 404, Message: "请求方法不存在"}, nil)
	//})
	//engine.Use(middlewares.CorsV2())
	engine.GET("/ping", func(c *gin.Context) {
		utils.Response(c,utils.HandlerSuccess,"word")
		})
	engine.GET("/getTask",controllers.GetTask.GetTask)
}