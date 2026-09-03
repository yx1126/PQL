package vo

import (
	"pql/DB/model"
	"strings"
)

type BaseAuth struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Icon         string `json:"icon"`
	Token        string `json:"token"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresTime  string `json:"expires_time"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type AuthVo struct {
	BaseAuth
	DriveId  int    `json:"id" gorm:"-"`
	Avatar   string `json:"avatar" gorm:"-"`
	Username string `json:"username" gorm:"-"`
	Nickname string `json:"nickname" gorm:"-"`
	VipType  int    `json:"vip_type" gorm:"-"`
	model.BaseModel
}

type SaveAuthVo struct {
	BaseAuth
}

func (ba *BaseAuth) IsAuth() bool {
	return strings.TrimSpace(ba.Token) != "" &&
		strings.TrimSpace(ba.RefreshToken) != "" &&
		ba.ExpiresIn != 0
}
