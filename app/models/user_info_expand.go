package models

import (
	"git.qutoutiao.net/todoList/app/entity/table"
)

type UserInfoExpandModel struct {
	*table.UserInfoExpandData
}

var UserInfoExpandTable = "user_info_expand"

func NewUserInfoExpandModel() *UserInfoExpandModel {
	return &UserInfoExpandModel{UserInfoExpandData: &table.UserInfoExpandData{}}
}

func (u *UserInfoExpandModel)FindByID(id int64) (err error) {
	err = Conn.Table(UserInfoExpandTable).Where("id = ?", id).Find(&u).Error
	return
}
////todo:这一部分是控制权限的层，或者一个叫数据返回控制层。
func (u *UserInfoExpandModel)GetPageSize()  int {
	err := u.FindByID(u.ID)
	if err != nil {
		return  0
	}
	return u.PageSize
}

