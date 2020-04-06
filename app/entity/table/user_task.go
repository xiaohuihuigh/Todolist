package table

type UserTask struct {
	ID int64 `gorm:"column:id"`
	TaskIDs string `gorm:"column:task_ids"`
}
