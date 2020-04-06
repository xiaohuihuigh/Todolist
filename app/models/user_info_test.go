package models_test

import (
	"fmt"
	"git.qutoutiao.net/todoList/app/models"
	"github.com/golib/assert"
	"testing"
)


func TestUserInfo(t *testing.T) {
	s:= models.NewUserInfoModel()
	sRes,err:=s.FindIn([]string{"1","3"})
	fmt.Println(len(sRes))
	it := assert.New(t)
	d:= models.NewUserInfoModel()
	d1:= models.NewUserInfoModel()
	d2:= models.NewUserInfoModel()
	d.ID = 2
	d.NickName = "test"
	d1.ID  = 2
	d1.NickName  = "test1"
	d2.ID  = 2
	d2.NickName  = "test1"
	//not record
	err = d.Update(d.ID)
	it.NotEmpty(err)
	err=d.Save()
	it.Empty(err)

	err=d.Save()
	it.Empty(err)
	//record id:2 nickname:test
	res:=models.NewUserInfoModel()
	err =res.FindByID(2)
	it.Empty(err)
	it.Equal(res.ID,d1.ID)

	//record id:2 nickname:test
	err=d1.Update(d1.ID)
	it.Empty(err)

	//record id:2 nickname:test1
	err =res.FindByID(d2.ID)
	it.Empty(err)
	it.Equal(res.ID ,d2.ID)
	it.Equal(res.NickName,d2.NickName)


	d3 := models.NewUserInfoModel()
	d3.ID = 3
	d3.NickName = "test2"
	err = d3.Update(2)
	it.Empty(err)

	d4 := models.NewUserInfoModel()
	d4.ID = 3
	d4.NickName = "test4"
	err = d4.Update(1)
	it.NotEmpty(err)

	//not record
	err=d1.Delete(3)
	it.Equal(err,nil)

	//not record
	err =res.FindByID(2)
	it.NotEmpty(err)

	_,err = res.Find()
	it.Empty(err)


}




