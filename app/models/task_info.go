package models

import (
	"git.qutoutiao.net/todoList/app/entity/table"
	"git.qutoutiao.net/todoList/app/utils"
	"github.com/jinzhu/gorm"
	"strconv"
	"sync"
	"time"
)

type TaskInfoModel struct {
	*table.TaskInfoData
	*gorm.DB
	IDs string
}

var TaskInfoTable = "task_info"

func NewTaskInfoModel() *TaskInfoModel {
	return &TaskInfoModel{TaskInfoData: &table.TaskInfoData{}, DB: Conn}
}

//todo:数据返回控制层


func (u *TaskInfoModel) SetTable() *gorm.DB {
	conn := u.Table(TaskInfoTable)
	return conn
}

func (u *TaskInfoModel) FindByID(id int64) (err error) {
	err = u.SetTable().Where("id = ?", id).Limit(1).Find(&u.TaskInfoData).Error
	return
}

func (u *TaskInfoModel) Find() (result []*table.TaskInfoData, err error) {
	err = u.SetTable().Debug().Find(&result).Error
	return
}
func (u *TaskInfoModel) FindByIDs(ids string) (result []*table.TaskInfoData, err error) {

	err = u.SetTable().Where("id in (?)", utils.String2Arr(ids)).Find(&result).Error
	return
}

func (u *TaskInfoModel) FindWhere() (result []*table.TaskInfoData, err error) {
	conn := u.SetTable()

	if u.ID != 0 {
		conn = conn.Where("id  = ?", u.ID)
	}
	if u.Title != "" {
		conn = conn.Where("title like ?", "%"+u.Title+"%")
	}
	if u.Context != "" {
		conn = conn.Where("context like ?","%"+ u.Context+"%")
	}
	if u.Type != 0 {
		conn = conn.Where("type = ?", u.Type)
	}
	if u.Priority != 0 {
		conn = conn.Where("priority = ?", u.Priority)
	}
	if u.ParentID != 0 {
		conn = conn.Where("parent_id = ?", u.ParentID)
	}
	err = conn.Debug().Find(&result).Error
	return
}
func (u *TaskInfoModel) FindIn(ids string, types string, priorities string, ) (result []*table.TaskInfoData, err error) {
	conn := u.SetTable()

	if ids != "" {
		conn = conn.Where("id in (?)", utils.String2Arr(ids))
	}

	if len(types) != 0 {
		conn = conn.Where("type in (?)", utils.String2Arr(types))
	}

	if len(priorities) != 0 {
		conn = conn.Where("priority in (?)", utils.String2Arr(priorities))
	}

	err = conn.Debug().Find(&result).Error
	return
}

func (u *TaskInfoModel) FindParent() (result []*table.TaskInfoData, err error) {
	err = u.SetTable().Where("id = ?", u.ParentID).Limit(1).Debug().Find(&result).Error
	return
}
func (u *TaskInfoModel) FindSub() (result []*table.TaskInfoData, err error) {
	err = u.SetTable().Where("id in (?)", utils.String2Arr(u.SubID)).Debug().Find(&result).Error
	return
}

func (u *TaskInfoModel) FindAll() (result []*table.TaskInfoData, err error) {

	err = u.SetTable().Find(&result).Error
	return
}

//todo:数据返回控制层
//func (u *TaskInfoModel) getIDs(dbData []*entity.TaskInfoData) (result string) {
//	result = ""
//	var resultArr []int64
//	for _, dbItem := range dbData {
//		resultArr = append(resultArr, dbItem.ID)
//	}
//	return utils.Int64Arr2String(resultArr)
//}

func (u *TaskInfoModel) Save() (err error) {
	uClone := NewTaskInfoModel()

	if uClone.FindByID(u.ID) != nil {
		err = u.Insert()
	} else {
		err = u.Update(u.ID)
	}
	return
}


func (u *TaskInfoModel) Insert() (err error) {
	u.Version = 1
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	err = u.SetTable().Debug().Create(&u.TaskInfoData).Error
	if u.SubID != "," {
		u.checkSubTaskParentID(u.SubID, ",")
	}
	if u.ParentID != -1 {
		u.checkParentTaskSubID(u.ParentID, -1)
	}
	return
}
func (u *TaskInfoModel) Update(id int64) (err error) {
	old := NewTaskInfoModel()
	err = old.FindByID(id)
	if err != nil {
		return
	}

	u.Version = old.Version + 1
	u.UpdatedAt = time.Now()
	err = u.SetTable().Where("id = ?", id).Debug().Update(&u.TaskInfoData).Error
	if u.SubID != old.SubID {
		u.checkSubTaskParentID(u.SubID, old.SubID)
	}
	if u.ParentID != old.ParentID {
		u.checkParentTaskSubID(u.ParentID, old.ParentID)
	}

	//wg.Wait()
	return
}

//todo 找出两个list不重合的不分old独有的删除parentid new独有的新增parentid，其他的不变。
func (u *TaskInfoModel) checkSubTaskParentID(newSubIDs string, oldSubIDs string) {
	delete,add := utils.MergeNewAndOld(newSubIDs,oldSubIDs)
	if delete != "," {
		dbRes, err := u.FindByIDs(delete)
		//todo:这里处理数据库查询错误，添加log？
		if err != nil {
			return
		}
		for _, dbItem := range dbRes {
			if dbItem.ParentID != -1 {
				dbItem.ParentID = -1
				newModel := NewTaskInfoModel()
				newModel.TaskInfoData = dbItem
				newModel.Save()
			}
		}
	}
	if add != "," {
		dbRes, err := u.FindByIDs(add)
		//todo:这里处理数据库查询错误，添加log？
		if err != nil {
			return
		}
		for _, dbItem := range dbRes {
			if dbItem.ParentID != u.ID {
				dbItem.ParentID = u.ID
				newModel := NewTaskInfoModel()
				newModel.TaskInfoData = dbItem
				newModel.Save()
			}
		}
	}

}

//todo 要从旧的父task中删除本task的id，然后在新的父task中添加
func (u *TaskInfoModel) checkParentTaskSubID(newParentID int64, oldParentID int64) {
	if oldParentID != -1 {
		newConn := NewTaskInfoModel()
		err := newConn.FindByID(oldParentID)
		if err != nil {
			return
		}
		if utils.ExistItem(utils.String2Arr(newConn.SubID), strconv.FormatInt(u.ID, 10)) {
			newConn.SubID = utils.DeleteItemInString(newConn.SubID, strconv.FormatInt(u.ID, 10))
			newConn.Save()
			return
		}
	}
	if newParentID != -1 {
		newConn := NewTaskInfoModel()
		err := newConn.FindByID(newParentID)
		if err != nil {
			return
		}
		if utils.ExistItem(utils.String2Arr(newConn.SubID), strconv.FormatInt(u.ID, 10)) {
			return
		}
		newConn.SubID = newConn.SubID + "," + strconv.FormatInt(u.ID, 10)
		newConn.Save()
	}

}
func (u *TaskInfoModel) Delete(id int64) (err error) {
	err = u.SetTable().Where("id = ?", id).Debug().Delete(struct{}{}).Error
	return
}
func (u *TaskInfoModel) DeleteByIDs(ids string) (err error) {
	wg := sync.WaitGroup{}

	for _, id := range utils.String2Arr(ids) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			idInt, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return
			}
			_ = u.Delete(idInt)

		}()
	}
	wg.Wait()
	return
}
func (t *TaskInfoModel) Filter(db *gorm.DB) {
	t.DB = db
}

func (t *TaskInfoModel) LimitByUser(db *gorm.DB) {
	ut := NewUserTaskModel()
	ut.DB = db
	err := ut.FindFirst()
	if err != nil {
		return
	}
	idArr := utils.String2Arr(ut.TaskIDs)
	t.DB = t.DB.Where("id in (?)", idArr)
	return
}

func (t *TaskInfoModel) LimitByTypeProject(db *gorm.DB) {

	taskType2Model := NewTaskForProjectModel()
	taskType2Model.DB = db
	dbRes, err := taskType2Model.Find()
	if err != nil {
		return
	}
	var idArr []int64
	for _, item := range dbRes {
		idArr = append(idArr, item.ID)
	}
	t.DB = t.DB.Where("id in (?)", idArr)
	return
}
