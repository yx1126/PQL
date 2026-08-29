package model

type Menu struct {
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Icon     string `json:"icon"`
	Path     string `json:"path"`
	Size     int    `json:"size" gorm:"default:16"`
	Hidden   int    `json:"hidden" gorm:"default:0"`
	BaseModel
}

func (*Menu) TableName() string {
	return "user_menu"
}
