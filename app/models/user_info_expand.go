package models

import "git.qutoutiao.net/todoList/app/entity"

type UserInfoExpandModel struct {
	*entity.UserInfoExpand
}

var UserInfoExpandTable = "user_info_expand"

func NewUserInfoExpandModel() *UserInfoExpandModel {
	return &UserInfoExpandModel{UserInfoExpand: &entity.UserInfoExpand{}}
}

func (u *UserInfoExpandModel) GetPageSize() (err error) {
	err = Conn.Table(UserInfoExpandTable).Where("id = ?", u.ID).Find(&u).Error
	return
}

func GetPageSize(id int64) (pageSize int) {
	userInfoExpandModel := NewUserInfoExpandModel()
	userInfoExpandModel.ID = id
	err := userInfoExpandModel.GetPageSize()
	if err != nil ||userInfoExpandModel.PageSize == 0 {
		userInfoExpandModel.ID = 1
		err=userInfoExpandModel.GetPageSize()
		if err != nil|| userInfoExpandModel.PageSize == 0{
			pageSize = 20
			return
		}
	}
	pageSize = userInfoExpandModel.PageSize
	return
}
