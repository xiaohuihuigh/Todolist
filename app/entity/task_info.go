package entity

type TaskInfo struct {
	ID        int64  `gorm:"column:id" json:"id"`
	Title     string `gorm:"column:title" json:"title"`
	Context   string `gorm:"column:context" json:"context"`
	Type      int    `gorm:"column:type" json:"type"`
	Priority  int    `gorm:"column:priority" json:"priority"`
	SubID     string `gorm:"column:sub_id" json:"sub_id"`
	ParentID  int64  `gorm:"column:parent_id" json:"parent_id"`
	Attention string `gorm:"column:attention" json:"attention"`
	Operator  int64  `gorm:"column:operator" json:"operator"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
	Version   int    `gorm:"column:version" json:"version"`
}
