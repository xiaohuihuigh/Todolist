package models

import (
	"todoList/app/entity/table"
	"todoList/app/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type UserInfoModel struct {
	*table.UserInfoData
	*gorm.DB
}

var UserInfoTable = "user_info"

func NewUserInfoModel() *UserInfoModel {
	return &UserInfoModel{UserInfoData: &table.UserInfoData{}, DB: Conn}
}
func (u *UserInfoModel) FindByID(id int64) (err error) {
	err = u.Table(UserInfoTable).Where("id = ?", id).Limit(1).Find(&u.UserInfoData).Error
	return
}
func (u *UserInfoModel) Find() (result []*table.UserInfoData, err error) {
	err = u.Table(UserInfoTable).Find(&result).Error
	return
}
func (u *UserInfoModel) FindIn(ids []string) (result []*table.UserInfoData, err error) {
	conn := u.Table(UserInfoTable)
	//if len(ids)!= 0{
	conn = conn.Where("id in (?)", ids)
	//}
	err = conn.Debug().Find(&result).Error
	return
}
func (u *UserInfoModel) Save() (err error) {
	uClone := u
	if uClone.FindByID(u.ID) != nil {
		err = u.Insert()
	} else {
		err = u.Update(u.ID)
	}
	return
}
func (u *UserInfoModel) Insert() (err error) {
	u.Version = 1
	u.CreatedAt = utils.TimeToString(time.Now())
	u.UpdatedAt = utils.TimeToString(time.Now())
	err = u.Table(UserInfoTable).Create(&u.UserInfoData).Error
	return
}
func (u *UserInfoModel) Update(id int64) (err error) {
	old := NewUserInfoModel()
	err = old.FindByID(id)
	if err != nil {
		return
	}
	u.Version = old.Version + 1
	u.UpdatedAt = utils.TimeToString(time.Now())
	err = u.Table(UserInfoTable).Where("id = ?", id).Debug().Update(&u.UserInfoData).Error
	return
}
func (u *UserInfoModel) Delete(id int64) (err error) {
	err = u.Table(UserInfoTable).Where("id = ?", id).Debug().Delete(struct {}{}).Error
	return
}
