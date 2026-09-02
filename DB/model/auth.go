package model

type Auth struct {
	Name         string `json:"name"`
	Type         string `json:"type" gorm:"uniqueIndex"`
	Icon         string `json:"icon"`
	Token        string `json:"token"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresTime  string `json:"expires_time"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	BaseModel
}

func (*Auth) TableName() string {
	return "user_auth"
}
