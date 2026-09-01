package model

type Auth struct {
	Type         string `json:"type" gorm:"uniqueIndex"`
	Token        string `json:"token"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresTime  string `json:"expires_time"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	BaseModel
}

func (*Auth) TableName() string {
	return "user_third_auth"
}
