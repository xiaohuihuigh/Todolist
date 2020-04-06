package table

import "time"

type TaskInfoData struct {
	ID            int64     `gorm:"column:id" json:"id"`
	Title         string    `gorm:"column:title" json:"title"`
	Context       string    `gorm:"column:context" json:"context"`
	Type          int       `gorm:"column:type" json:"type"`
	Priority      int       `gorm:"column:priority" json:"priority"`
	SubID         string    `gorm:"column:sub_id" json:"sub_id"`
	ParentID      int64     `gorm:"column:parent_id" json:"parent_id"`
	Operator      int64     `gorm:"column:operator" json:"operator"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	Version       int       `gorm:"column:version" json:"version"`
	SubTask []*TaskInfoData
	ProjectEx *TaskInfoForProjectData

}

//id         integer PRIMARY KEY AUTOINCREMENT, --唯一标示
//title      text     NOT NULL,                 --标题
//context    text    DEFAULT NULL,              --内容
//type       integer DEFAULT 1,                 --task的类型 1,2 todo|project
//priority   integer DEFAULT 1,                 --优先级 1-9  映射列表
//sub_id     text    DEFAULT ',',               --子task的所有id用`,`隔开
//parent_id  integer default -1,                --父task的id
//operator   integer  NOT NULL,                 -- 操作人
//created_at datetime NOT NULL,
//updated_at datetime NOT NULL,
//version    integer  NOT NULL