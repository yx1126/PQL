package model

type Live struct {
	RoomId    string `json:"roomId"`
	Type      string `json:"type"`
	Sort      *int   `json:"sort"`
	IsSpecial int    `json:"isSpecial" gorm:"default:0"`
	BaseModel
}

func (*Live) TableName() string {
	return "user_live"
}
