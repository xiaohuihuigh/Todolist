package controllers

import (
	"git.qutoutiao.net/todoList/app/entity/table"
	"git.qutoutiao.net/todoList/app/utils"
	"github.com/gin-gonic/gin"
)

var AddTask *addTask
type addTask struct {}

type AddTaskRequest struct {
	Operator int64
	Task table.TaskInfoData
}
func (t *addTask) AddTask(ctx *gin.Context){
	//request:= parseAddTaskRequest(ctx)
	//if request.Operator == 0{
	//	utils.Response(ctx,utils.ErrorApiHandler,"未登录")
	//	return
	//}
	//taskResult :=services.NewTaskResult(request.Operator)
	//services.TaskResult.GetAllTask()
}
func parseAddTaskRequest(ctx *gin.Context)*AddTaskRequest{
	request := &AddTaskRequest{}
	request.Operator = utils.GetStr2Int64(ctx.PostForm("operator"),0)
	request.Task.Title = ctx.PostForm("title")
	request.Task.Context = ctx.PostForm("context")
	request.Task.Type = utils.GetStr2Int(ctx.PostForm("type"),1)
	//...
	return request
}