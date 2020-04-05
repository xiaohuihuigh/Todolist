package models_test

import (
	"git.qutoutiao.net/todoList/app/models"
	"github.com/golib/assert"
	"strconv"
	"testing"
)

func TestTaskInfo(t *testing.T){
	task1 := models.NewTaskInfoModel()
	task1.Title = "测试用title"
	task1.Context = "测试用context"
	err:=task1.Save()
	it := assert.New(t)
	it.Empty(err)

	task2 := models.NewTaskInfoModel()
	task2.Title = "测试用例2"
	task2.SubID = strconv.FormatInt(task1.ID,10)
	task2.Save()
	task1.FindByID(task1.ID)
	it.NotEmpty(task1.ParentID)

	task3 := models.NewTaskInfoModel()
	task3.Title = "测试用例3"
	task3.ParentID = task2.ID
	task3.Save()
	task1.FindByID(task1.ID)
	it.NotEmpty(task3.ParentID)

	task2.SubID = ""
	task2.Save()
	task1.FindByID(task1.ID)
	task3.FindByID(task3.ID)
	it.Equal(task1.ParentID,int64(-1))
	it.Equal(task3.ParentID,int64(-1))

}