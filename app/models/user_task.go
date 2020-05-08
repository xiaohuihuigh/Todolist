package models

import (
	"todoList/app/entity/table"
	"github.com/jinzhu/gorm"
)

type UserTaskModel struct {
	*table.UserTask
	*gorm.DB
}

var UserTaskTable = "user_task"

func NewUserTaskModel() *UserTaskModel {
	return &UserTaskModel{UserTask: &table.UserTask{}, DB: Conn}
}

func (ut *UserTaskModel) FindByID(id int) (err error) {
	err = ut.Table(UserTaskTable).Where("id = ?", id).First(&ut).Error
	//log
	return
}
func (ut *UserTaskModel) FindFirst() (err error) {
	err = ut.Table(UserTaskTable).First(&ut).Error
	return
}
