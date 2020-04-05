package entity

type UserInfoExpand struct {
	ID int64 `gorm:"column:id"`
	Avatar string `gorm:"column:avatar"`
	PageSize int `gorm:"column:page_size"`
	WeixinNickname string `gorm:"column:weixin_nickname"`
	WeixinAvatar string `gorm:"column:weixin_avatar"`
	ContinueLoginDay int `gorm:"column:continue_login_day"`
	DepartmentName string `gorm:"column:departmentName"`
	Position string `gorm:"column:position"`
	LastLoginDate string `gorm:"last_login_date"`
}
