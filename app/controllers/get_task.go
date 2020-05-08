package controllers

import (
	"todoList/app/entity/service"
	"todoList/app/services"
	"todoList/app/utils"
	"github.com/gin-gonic/gin"
)

var GetTask *getTask
type getTask struct {}

//type GetTaskRequest struct {
//	Operator int64
//}
func (gt *getTask) GetTask(ctx *gin.Context){
	search:=parseGetTaskRequest(ctx)
	if search.Operator == 0{
		utils.Response(ctx,utils.ErrorApiHandler,"not login")
		return
	}
	result := services.NewTaskResult(search.Operator)
	result.Search = search
	err := result.GetResult()
	if err != nil{
		utils.Response(ctx,utils.ErrorApiHandler,"get task error")
		return
	}
	utils.Response(ctx,utils.HandlerSuccess,result)
}
func parseGetTaskRequest(ctx *gin.Context)*service.Search{
	request := service.NewSearch()
	request.Operator =utils.GetStr2Int64(ctx.Query("operator"),0)
	request.Type = utils.GetStr2Int(ctx.Query("type"),0)
	request.IDs = ctx.Query("ids")
	return request
}