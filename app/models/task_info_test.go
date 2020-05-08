package models_test

import (
	"todoList/app/entity/table"
	"todoList/app/models"
	"todoList/app/utils"
	"github.com/golib/assert"
	"strconv"
	"testing"
)

func TestTaskInfo(t *testing.T){
	it := assert.New(t)


	task1 := models.NewTaskInfoModel()
	task1.Title = "测试用title"
	task1.Context = "测试用context"
	err:=task1.Save()
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

	task3.SubID = strconv.FormatInt(task2.ID,10)
	task3.Save()

	task2.FindByID(task2.ID)
	it.Equal(task2.ParentID,task3.ID)

	task3.SubID = strconv.FormatInt(task1.ID,10)
	task3.Save()

	task2.FindByID(task2.ID)
	it.Equal(task2.ParentID,int64(-1))
	task1.FindByID(task1.ID)
	it.Equal(task1.ParentID,task3.ID)

}
func TestFindWhere(t *testing.T){
	it := assert.New(t)
	var err error
	var dbRes []*table.TaskInfoData
	task1 := models.NewTaskInfoModel()
	task1.Title = "测试findwhere用title1"
	task1.Context = "测试findwhere用context1"
	task1.Type = 1
	task1.Priority = 1
	err=task1.Update(task1.ID)
	it.NotEmpty(err)
	err=task1.Save()
	it.Empty(err)

	task2 := models.NewTaskInfoModel()
	task2.Title = "测试findwhere用title2"
	task2.Context = "测试findwhere用context2"
	task2.Type = 2
	task2.Priority = 2
	task2.ParentID = task1.ID
	err=task2.Save()
	it.Empty(err)

	task1.FindByID(task1.ID)
	task2.FindByID(task2.ID)

	find := models.NewTaskInfoModel()
	dbRes,err=find.FindAll()
	it.Empty(err)
	//it.Equal(len(dbRes),2)


	find1 := models.NewTaskInfoModel()
	find1.Title = "测试findwhere用title"
	find1.Context = "测试findwhere用context"
	dbRes,err=find1.FindWhere()
	it.Equal(len(dbRes),2)

	find2 := models.NewTaskInfoModel()
	find2.Title = "测试findwhere用title"
	find2.Context = "测试findwhere用context"
	find2.ParentID = task1.ID
	find2.Type = task2.Type
	find2.Priority = task2.Type
	dbRes,err=find2.FindWhere()
	it.Equal(len(dbRes),1)

	find3 := models.NewTaskInfoModel()
	find3.Title = "测试findwhere用title"
	find3.Context = "测试findwhere用context"
	find3.ParentID = task1.ID
	find3.Type = task2.Type
	//find3.Priority = task2.Type
	//find3.Attention = task2.Attention
	find3.CreatedAt = task2.CreatedAt
	find3.ID = task1.ID
	dbRes,err=find3.FindWhere()
	it.Equal(len(dbRes),0)

	find4 := models.NewTaskInfoModel()
	dbRes,err=find4.FindIn("","1,2","1,2")
	it.Equal(len(dbRes),2)


	find5 := models.NewTaskInfoModel()
	dbRes,err=find5.FindIn(utils.Int64Arr2String([]int64{task1.ID,task2.ID}),"","")
	it.Equal(len(dbRes),2)

	find6 := task1
	dbRes,err=find6.FindSub()
	it.Equal(len(dbRes),1)

	find7:= task2
	dbRes,err=find7.FindParent()
	it.Equal(len(dbRes),1)

	task1.Delete(task1.ID)
	task2.DeleteByIDs(utils.Int64Arr2String([]int64{task1.ID,task2.ID}))
}