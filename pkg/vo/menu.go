package vo

import "pql/DB/model"

type BaseMenuVo struct {
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Icon     string `json:"icon"`
	Path     string `json:"path"`
	Size     int    `json:"size" gorm:"default:16"`
	Hidden   int    `json:"hidden"`
}

type MenuVo struct {
	BaseMenuVo
	model.BaseModel
}

type UpdateMenuVo struct {
	BaseMenuVo
	BaseVo
}
