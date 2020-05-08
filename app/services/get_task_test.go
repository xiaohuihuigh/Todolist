package services_test

import (
	"todoList/app/models"
	"todoList/app/services"
	"github.com/golib/assert"
	"testing"
)

func TestMain(m *testing.M) {
	err := models.Setup()
	it := assert.Assertions{}
	it.Empty(err)
	m.Run()
}
func TestGetTask(t *testing.T) {
	// 查询所有的task
	result := services.NewTaskResult(2)
	result.Search.PageNum = 1
	result.Search.Type = 1
	err := result.GetResult()
	it := assert.New(t)
	it.Empty(err)
	it.Equal(len(result.Task), 2)

	//查询所有的todo
	result = services.NewTaskResult(1)
	result.Search.Type = 1
	err = result.GetResult()
	it.Empty(err)
	it.Equal(len(result.Task), 2)

	//查询所有的project
	result = services.NewTaskResult(1)
	result.Search.Type = 2
	//result.Search.Process = 44
	err = result.GetResult()
	it.Empty(err)
	it.Equal(len(result.Task), 3)

	//查询一条todo，by id
	result = services.NewTaskResult(1)
	result.Search.Type = 1
	result.Search.TaskSearch.ID = 4
	err = result.GetResult()
	it.Empty(err)
	it.Equal(len(result.Task), 1)
	it.Equal(result.Task[0].ID, int64(4))

	//查询一条project，by id
	result = services.NewTaskResult(1)
	result.Search.Type = 2
	result.Search.TaskType2Search.ID = 1
	//result.Search.Process = 44
	err = result.GetResult()
	it.Empty(err)
	it.Equal(len(result.Task), 1)
	it.Equal(result.Task[0].ID, int64(1))

	//通过筛选条件，查询todo
	result = services.NewTaskResult(1)
	result.Search.Type = 1
	result.Search.TaskSearch.Title = "todo"
	//result.Search.Process = 44
	err = result.GetResult()
	it.Empty(err)
	it.Equal(len(result.Task), 2)
	//it.Equal(result.Task[0].ID, 4)

	//通过筛选条件，查询project
	result = services.NewTaskResult(1)
	result.Search.Type = 2
	result.Search.TaskSearch.Title = "project"
	result.Search.TaskType2Search.Progress = 44
	//result.Search.Process = 44
	err = result.GetResult()
	it.Empty(err)
	it.Equal(len(result.Task), 1)
	it.Equal(result.Task[0].ID, int64(1))

	//通过筛选条件，查询task
	result = services.NewTaskResult(1)
	result.Search.IDs = "1,2,3,4"
	//result.Search.TaskType2Search.Process = 44
	//result.Search.Process = 44
	err = result.GetResult()
	it.Empty(err)
	it.Equal(len(result.Task), 4)
	//it.Equal(result.Task[0].ID, 1)
}
