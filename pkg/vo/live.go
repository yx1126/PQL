package vo

import (
	"errors"
	"pql/DB/model"
	"pql/pkg/parse/live"
	"strings"
)

type LiveParams struct {
	Type      string `json:"type"`
	IsSpecial string `json:"isSpecial"`
}

type LiveVo struct {
	live.RoomInfo `gorm:"-"`
	RoomId        string `json:"roomId"`
	Rid           string `json:"rid" gorm:"-"`
	Type          string `json:"type"`
	Sort          *int   `json:"sort"`
	IsSpecial     int    `json:"isSpecial"`
	model.BaseModel
}

type CreateLiveVo struct {
	RoomId    string `json:"roomId"`
	Type      string `json:"type"`
	IsSpecial int    `json:"isSpecial"`
}

type LiveSearchVo struct {
	Page    string `json:"page"`
	Keyword string `json:"keyword"`
	Type    string `json:"type"`
}

type UpdateLiveVo struct {
	Sort      *int `json:"sort"`
	IsSpecial int  `json:"isSpecial"`
	BaseVo
}

func (l *CreateLiveVo) Valid() error {
	if strings.TrimSpace(l.RoomId) == "" || strings.TrimSpace(l.Type) == "" {
		return errors.New("房间号或类型不能为空")
	}
	return nil
}
