package services

import (
	"git.qutoutiao.net/todoList/app/entity/service"
	"git.qutoutiao.net/todoList/app/entity/table"
	"git.qutoutiao.net/todoList/app/models"
	"git.qutoutiao.net/todoList/app/utils"
	"github.com/jinzhu/gorm"
)

//var TaskResult *TaskResult
type TaskResult struct {
	*service.TaskResult
}

func NewTaskResult(userID int64) *TaskResult {
	return &TaskResult{TaskResult: service.NewTaskResult(userID),}
}

func (tr *TaskResult) GetResult() (err error) {
	err = tr.GetAllTask()
	if err != nil {
		//log
	}
	err = tr.FilterRootTask()
	err = tr.GetOperatorInfo()
	if err != nil {
		//log
	}
	return
}
func (tr *TaskResult) FilterRootTask() (err error) {
	taskMap:= map[int64]*table.TaskInfoData{}
	for _, item := range tr.Task {
		//if item.Type == 2{
			taskMap[item.ID] = item
		//}
	}
	for _, v := range taskMap {
		if _,ok:= taskMap[v.ParentID];!ok {
			tr.RootTask = append(tr.RootTask, v)
		}
	}
	for _, item := range tr.RootTask {
		tr.FilterSubTask(item,taskMap)
	}
	return
}
func (tr *TaskResult) FilterSubTask(item *table.TaskInfoData,taskMap map[int64]*table.TaskInfoData) {
	if item.SubID == "," ||item.Type == 1 {
		return
	}
	for _, i := range utils.String2Int64Arr(item.SubID){
		if _,ok := taskMap[i];ok {
			item.SubTask = append(item.SubTask, taskMap[i])
			if taskMap[i].Type == 2 {
				tr.FilterSubTask(taskMap[i],taskMap)
			}
		}
	}
}

func (tr *TaskResult) GetOperatorInfo() (err error) {
	UserModel := models.NewUserInfoModel()
	if tr.Search.Operator != 0 {
		err = UserModel.FindByID(tr.Search.Operator)
		tr.Operator = UserModel.UserInfoData
	}
	return
}
func (tr *TaskResult) GetAllTask() (err error) {
	err = tr.GetTask()

	//if tr.Search.Type == 0 {
	//	err = tr.GetProject()
	//	err = tr.GetTodo()
	//} else if tr.Search.Type == 2 {
	//	err = tr.GetProject()
	//} else if tr.Search.Type == 1 {
	//	err = tr.GetTodo()
	//}
	if err != nil {
		return
	}
	for _,item := range tr.Task{
		if item.Type == 1{
			tr.Todos = append(tr.Todos,item)
		}else if item.Type == 2{
			taskForProjectModel := models.NewTaskForProjectModel()
			_ = taskForProjectModel.FindByID(item.ID)
			item.ProjectEx = taskForProjectModel.TaskInfoForProjectData
			tr.Projects = append(tr.Projects,item)
		}
	}
	//for _, item := range tr.Projects {
	//
	//
	//}

	//for _, item := range tr.Todos {
	//	taskType2Model := my_gorm.NewTaskType1Model()
	//	_ = taskType2Model.FindByID(item.ID)
	//	item.ProjectEx = taskType2Model.TaskType1
	//
	//}
	//tr.Task = append(tr.Task, tr.Projects...)
	//tr.Task = append(tr.Task, tr.Todos...)

	return
}

//func (tr *TaskResult) GetExAndSubTask(item *entry.Task) {
//	taskType2Model := my_gorm.NewTaskForProjectModel()
//	_ = taskType2Model.FindByID(item.ID)
//	item.ProjectEx = taskType2Model.TaskType2
//	if item.SubID == "," {
//		return
//	}
//	result, _ := tr.GetSubProject(utils.String2Arr(item.SubID))
//	item.SubTask = result
//	for _, i := range item.SubTask {
//		tr.GetExAndSubTask(i)
//	}
//}

//func (tr *TaskResult) GetProject() (err error) {
//	tr.Search.Type = 2
//	taskModel := models.NewTaskInfoModel()
//	taskModel.LimitByUser(tr.SearchFilterUserTask())
//	taskModel.LimitByTypeProject(tr.SearchFilterTaskType2())
//	taskModel.Filter(tr.SetOffsetAndLimit(taskModel.DB))
//	taskModel.Filter(tr.SearchFilterTask(taskModel.DB))
//	tr.Projects, err = taskModel.Find()
//	//tr.Task = append(tr.Task,tr.Projects...)
//	return
//}

//func (tr *TaskResult) GetSubProject(ids []string) (result []*entry.Task, err error) {
//	taskModel := my_gorm.NewTaskModel()
//	taskModel.LimitByUser(tr.SearchFilterUserTask())
//	taskModel.LimitByTypeProject(tr.SearchFilterTaskType2())
//	taskModel.Conn = tr.SearchFilterTask(taskModel.Conn)
//	result, err = taskModel.FindByIDs(ids)
//	return
//}

//func (tr *TaskResult) GetTodo() (err error) {
//	taskModel := models.NewTaskInfoModel()
//	tr.Search.Type = 1
//	taskModel.LimitByUser(tr.SearchFilterUserTask())
//	//taskModel.LimitByTypeProject(tr.SearchFilterTaskType2())
//	taskModel.Filter(tr.SetOffsetAndLimit(taskModel.DB))
//	taskModel.Filter(tr.SearchFilterTask(taskModel.DB))
//	tr.Todos, err = taskModel.Find()
//	//tr.Task = append(tr.Task,tr.Todos...)
//	return
//}

func (tr *TaskResult) GetTask() (err error) {
	taskModel := models.NewTaskInfoModel()
	taskModel.LimitByUser(tr.SearchFilterUserTask())
	//TaskInfoForProjectData 要进行判断
	//if tr.Search.Type ==1{
	//	taskModel.LimitByTypeProject(tr.SearchFilterTaskType1())
	//}
	if tr.Search.Type ==2{
		taskModel.LimitByTypeProject(tr.SearchFilterTaskType2())
	}
	//taskModel.Filter(tr.SetOffsetAndLimit(taskModel.DB))
	taskModel.Filter(tr.SearchFilterTask(taskModel.DB))

	tr.Task, err = taskModel.Find()
	//tr.Task = append(tr.Task,tr.Todos...)
	return
}



func (tr *TaskResult) SearchFilterTaskType2() *gorm.DB {
	conn := models.Conn
	if tr.Search.TaskType2Search.ID != 0 {
		conn = conn.Where("id = ?", tr.Search.TaskType2Search.ID)
	}
	//else search
	if tr.Search.Progress != 0 {
		conn = conn.Where("progress = ?", tr.Search.Progress)
	}
	return conn
}

func (tr *TaskResult) SearchFilterTask(conn *gorm.DB) *gorm.DB {
	if tr.Search.Type != 0 {
		conn = conn.Where("type = ?", tr.Search.Type)
	}
	if tr.Search.IDs != "" {
		conn = conn.Where("id in (?)", utils.String2Arr(tr.Search.IDs))
	}
	if tr.Search.TaskSearch.ID != 0 {
		conn = conn.Where("id = ?", tr.Search.TaskSearch.ID)
	}
	if tr.Search.TaskSearch.Title != "" {
		conn = conn.Where("title like (?)", "%"+tr.Search.TaskSearch.Title+"%")
	}
	return conn
}
func (tr *TaskResult) SearchFilterUserTask() *gorm.DB {
	conn := models.Conn
	if tr.Search.Operator != 0 {
		conn = conn.Where("id = ?", tr.Search.Operator)
	}
	return conn
}
func (tr *TaskResult) SetOffsetAndLimit(conn *gorm.DB) *gorm.DB {
	//conn := u.SetIDs()
	pageSize := tr.GetPageSize()
	if tr.Search.PageNum == 0 {
		tr.Search.PageNum = 1
	}
	conn = conn.Offset(pageSize * (tr.Search.PageNum - 1)).Limit(pageSize)
	//SetOffsetAndLimit(conn, operator, pageNum)
	return conn
}
func (tr *TaskResult) SetIDs(conn *gorm.DB) *gorm.DB {
	conn = conn.Where("id in (?)", utils.String2Arr(tr.Search.IDs))
	return conn
}

//这一部分是控制权限的层，或者一个叫数据返回控制层。
func (tr *TaskResult) GetPageSize() (pageSize int) {
	userInfoExpandModel := models.NewUserInfoExpandModel()
	userInfoExpandModel.ID = tr.Operator.ID
	pageSize = userInfoExpandModel.GetPageSize()
	if pageSize == 0 {
		userInfoExpandModel.ID = 1
		pageSize = userInfoExpandModel.GetPageSize()
		if pageSize == 0 {
			pageSize = 20
			return
		}
	}
	return
}
