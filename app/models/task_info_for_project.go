package models

import (
	"git.qutoutiao.net/todoList/app/entity/table"
	"github.com/jinzhu/gorm"
)

type TaskInfoFroProjectModel struct {
	*table.TaskInfoForProjectData
	*gorm.DB
}
var TaskInfoExpandProject =  "task_info_for_project"
func NewTaskForProjectModel() *TaskInfoFroProjectModel {
	return &TaskInfoFroProjectModel{TaskInfoForProjectData: &table.TaskInfoForProjectData{}, DB: Conn}
}
func (tt2 *TaskInfoFroProjectModel) Find() (result []*table.TaskInfoForProjectData, err error) {
	err = tt2.Table(TaskInfoExpandProject).Debug().Find(&result).Error
	return
}
func (tt2 *TaskInfoFroProjectModel) FindByID(id int64) (err error) {
	err = tt2.Table(TaskInfoExpandProject).Where("id = ?", id).Debug().First(&tt2).Error
	return
}
