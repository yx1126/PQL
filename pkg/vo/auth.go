package vo

import "pql/DB/model"

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
	model.BaseModel
}

type SaveAuthVo struct {
	BaseAuth
}
