package model

import (
	"pql/pkg/utils/types"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt types.Datetime `json:"createdAt"`
	UpdatedAt types.Datetime `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
