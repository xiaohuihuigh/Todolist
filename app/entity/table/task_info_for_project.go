package table

import "time"

type TaskInfoForProjectData struct {
	ID            int64     `gorm:"column:id" json:"id"`
	TaskStartTime time.Time `gorm:"column:task_start_time" json:"task_start_time"`
	EndTime       time.Time `gorm:"column:end_time" json:"end_time"`
	PlanEndTime   time.Time `gorm:"column:plan_end_time" json:"plan_end_time"`
	ActualEndTime time.Time `gorm:"column:actual_end_time" json:"actual_end_time"`
	Progress      int       `gorm:"progress"`
	Status        int       `gorm:"column:status"`
	Operator      int64     `gorm:"column:operator" json:"operator"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	Version       int       `gorm:"column:version" json:"version"`
}
