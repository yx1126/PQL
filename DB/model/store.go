package model

type Store struct {
	Key   string `json:"key" gorm:"uniqueIndex"`
	Value string `json:"value"`
	BaseModel
}

func (*Store) TableName() string {
	return "user_store"
}
