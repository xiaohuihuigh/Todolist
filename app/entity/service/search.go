package service

import "git.qutoutiao.net/todoList/app/entity/table"

type Search struct {
	Operator         int64 `json:"operator_id"`
	*TaskSearch      `json:"task_search"`
	*TaskType2Search `json:"task_type2_search"`
	PageNum          int    `json:"page_num"`
	Type             int    `json:"type"`
	IDs              string `json:"ids"`
}

type TaskSearch struct {
	*table.TaskInfoData
}

type TaskType2Search struct {
	*table.TaskInfoForProjectData
}

func NewSearch() *Search {
	return &Search{
		Operator:        0,
		TaskSearch:      &TaskSearch{TaskInfoData: &table.TaskInfoData{}},
		TaskType2Search: &TaskType2Search{TaskInfoForProjectData: &table.TaskInfoForProjectData{}},
		PageNum:         1,
		Type:            0,
		IDs:             "",
	}
}
