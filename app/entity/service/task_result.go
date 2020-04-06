package service

import (
	"git.qutoutiao.net/todoList/app/entity/table"
)

type TaskResult struct{
	Task     []*table.TaskInfoData `json:"task"`
	Todos    []*table.TaskInfoData `json:"todo"`
	Projects []*table.TaskInfoData `json:"project"`
	RootTask []*table.TaskInfoData `json:"root_task"`
	Operator *table.UserInfoData   `json:"operator"`
	Search   *Search               `json:"search"`
}
func NewTaskResult (userID int64)*TaskResult{
	return &TaskResult{
		Operator: &table.UserInfoData{ID: userID},
		Search:   NewSearch(),
	}
}
