package table

type UserInfoData struct {
	ID int64 `gorm:"column:id"`
	NickName string `gorm:"column:nickname"`
	Phone string `gorm:"column:phone"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
	Version int `gorm:"column:version"`
}
