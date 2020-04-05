package models

import (
	"git.qutoutiao.net/todoList/app/entity"
	"git.qutoutiao.net/todoList/app/utils"
	"github.com/jinzhu/gorm"
	"strconv"
	"sync"
	"time"
)

type TaskInfoModel struct {
	*entity.TaskInfo
	*gorm.DB
	IDs string
}

var TaskInfoTable = "task_info"

func NewTaskInfoModel() *TaskInfoModel {
	return &TaskInfoModel{TaskInfo: &entity.TaskInfo{}, DB: Conn}
}

//避免重复代码
func (u *TaskInfoModel) SetIDsOffsetAndLimit(operator int64, pageNum int) *gorm.DB {
	conn := u.SetIDs()
	SetOffsetAndLimit(conn, operator, pageNum)
	return conn
}
func (u *TaskInfoModel) SetIDs() *gorm.DB {
	conn := u.Where("id in (?)", utils.String2Arr(u.IDs))
	return conn
}

func (u *TaskInfoModel) SetTable() *gorm.DB {
	conn := u.Table(TaskInfoTable)
	return conn
}

func (u *TaskInfoModel) FindByID(id int64) (err error) {
	err = u.SetTable().Where("id = ?", id).Limit(1).Find(&u.TaskInfo).Error
	return
}

func (u *TaskInfoModel) FindByIDs(ids string) (result []*entity.TaskInfo, err error) {

	err = u.SetTable().Where("id in (?)", utils.String2Arr(ids)).Find(&result).Error
	return
}

func (u *TaskInfoModel) FindWhere() (result []*entity.TaskInfo, err error) {
	conn := u.SetTable()

	if u.ID != 0 {
		conn = conn.Where("id  = ?", u.ID)
	}
	if u.Title != "" {
		conn = conn.Where("title like '%?%'", u.Title)
	}
	if u.Context != "" {
		conn = conn.Where("context like '%?%", u.Context)
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
	if u.Attention != "" {
		conn = conn.Where("attention like '%?%'", u.Attention)
	}
	if u.CreatedAt != "" {
		conn = conn.Where("created_at like '%?%'", u.CreatedAt)
	}
	err = conn.Find(&result).Error
	return
}
func (u *TaskInfoModel) FindIn(ids string, types string, priorities string, ) (result []*entity.TaskInfo, err error) {
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

	err = conn.Find(&result).Error
	return
}

func (u *TaskInfoModel) FindParent() (result []*entity.TaskInfo, err error) {
	err = u.SetTable().Where("id = ?", u.ParentID).Limit(1).Find(&result).Error
	return
}
func (u *TaskInfoModel) FindSub(pageNum int) (result []*entity.TaskInfo, err error) {
	err = u.SetTable().Where("id in (?)", utils.String2Arr(u.SubID)).Find(&result).Error
	return
}

func (u *TaskInfoModel) FindAll() (result []*entity.TaskInfo, err error) {

	err = u.SetTable().Find(&result).Error
	return
}

func (u *TaskInfoModel) getIDs(dbData []*entity.TaskInfo) (result string) {
	result = ""
	var resultArr []int64
	for _, dbItem := range dbData {
		resultArr = append(resultArr, dbItem.ID)
	}
	return utils.Int64Arr2String(resultArr)
}

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
	u.CreatedAt = utils.TimeToString(time.Now())
	u.UpdatedAt = utils.TimeToString(time.Now())
	err = u.SetTable().Debug().Create(&u.TaskInfo).Error
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
	u.UpdatedAt = utils.TimeToString(time.Now())
	err = u.SetTable().Where("id = ?", id).Debug().Update(&u.TaskInfo).Error
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
		if err != nil {
			return
		}
		for _, dbItem := range dbRes {
			if dbItem.ParentID != -1 {
				dbItem.ParentID = -1
				newModel := NewTaskInfoModel()
				newModel.TaskInfo = dbItem
				newModel.Save()
			}
		}
	}
	if add != "," {
		dbRes, err := u.FindByIDs(add)
		if err != nil {
			return
		}
		for _, dbItem := range dbRes {
			if dbItem.ParentID != u.ID {
				dbItem.ParentID = u.ID
				newModel := NewTaskInfoModel()
				newModel.TaskInfo = dbItem
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
	err = u.SetTable().Where("id = ?", id).Delete(struct{}{}).Error
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
