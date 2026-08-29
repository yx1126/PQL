package live

import (
	"errors"
	"pql/pkg/request"
	"strconv"
)

const (
	UnknowErr = "未知错误！"
	TypeErr   = "类型不存在！"
	RoomErr   = "房间不存在！"
	ParserErr = "解析失败！"
	LineErr   = "线路不存在！"
)

type RoomInfo struct {
	RoomId      string `json:"roomId"`
	Type        string `json:"type"`
	RoomName    string `json:"roomName"`
	OwnerName   string `json:"ownerName"`
	OwnerAvatar string `json:"ownerAvatar"`
	IsLive      bool   `json:"isLive"`
	IsReplay    bool   `json:"isReplay"`
	RoomPic     string `json:"roomPic"`
	Description string `json:"description"`
}

type LiveInfo struct {
	RoomId       string         `json:"roomId"`
	Url          string         `json:"url"`
	VideoType    string         `json:"videoType"`
	RtmpCdn      string         `json:"rtmpCdn"`
	Rate         string         `json:"rate"`
	Multirates   []Multirate    `json:"multirates"`
	CdnsWithName []CdnsWithName `json:"cdnsWithName"`
}

type Multirate struct {
	Name string `json:"name"`
	Rate string `json:"rate"`
}

type CdnsWithName struct {
	Name string `json:"name"`
	Cdn  string `json:"cdn"`
	// huya,douyin
	StreamName string `json:"streamName"`
}

type LiveParams struct {
	Page    string `json:"page"`
	Keyword string `json:"keyword"`
	Type    string `json:"type"`
}

type Live struct {
	douyu  *Douyu
	huya   *Huya
	douyin *Douyin
	ttwid  string
}

type LiveResponse[T any] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
}

func New(http *request.Http) *Live {
	return &Live{
		douyu:  NewDouyu(http),
		huya:   NewHuya(http),
		douyin: NewDouyin(http),
	}
}

func (l *Live) SetTtwid(value string) {
	l.ttwid = value
}

func (l *Live) GetRoomWebUri(roomId, roomType string) string {
	switch roomType {
	case "1":
		return l.douyu.GetRoomWebUri(roomId)
	case "2":
		return l.huya.GetRoomWebUri(roomId)
	case "3":
		return l.douyin.GetRoomWebUri(roomId)
	default:
		return ""
	}
}

func (l *Live) GetInfo(roomId, roomType string) (*RoomInfo, error) {
	var v *RoomInfo
	var err error
	switch roomType {
	case "1":
		v, err = l.douyu.GetInfo(roomId)
	case "2":
		v, err = l.huya.GetInfo(roomId)
	case "3":
		v, err = l.douyin.GetInfo(roomId, l.ttwid)
	default:
		return nil, errors.New(TypeErr)
	}
	if err != nil {
		return nil, err
	}
	v.Type = roomType
	return v, nil
}

func (l *Live) GetPlayUrl(roomId, roomType string, extra map[string]string) (*LiveInfo, error) {
	switch roomType {
	case "1":
		return l.douyu.GetPlayUrl(roomId, extra)
	case "2":
		return l.huya.GetPlayUrl(roomId, extra)
	case "3":
		return l.douyin.GetPlayUrl(roomId, extra)
	default:
		return nil, errors.New(TypeErr)
	}
}

func (l *Live) Search(params LiveParams) (*LiveResponse[RoomInfo], error) {
	switch params.Type {
	case "1":
		res, err := l.douyu.Search(params)

		if err != nil {
			return nil, err
		}

		var result LiveResponse[RoomInfo]

		for _, v := range res.RelateShow {
			result.Data = append(result.Data, RoomInfo{
				RoomId:      strconv.Itoa(v.Rid),
				Type:        params.Type,
				RoomName:    v.RoomName,
				OwnerName:   v.NickName,
				OwnerAvatar: v.Avatar,
				IsLive:      true,
				RoomPic:     v.RoomSrc,
			})
		}
		result.Total = res.Total

		return &result, nil

	case "2":
		res, err := l.huya.Search(params)
		if err != nil {
			return nil, err
		}
		var result LiveResponse[RoomInfo]

		// 处理虎牙分页时第一页往后返回前页的数据
		if len(res.Docs) > 20 {
			page, err := strconv.Atoi(params.Page)
			if err != nil {
				return nil, err
			}
			size := (page - 1) * 20
			if size < len(res.Docs) {
				res.Docs = res.Docs[size:]
			}
		}

		for _, v := range res.Docs {
			result.Data = append(result.Data, RoomInfo{
				RoomId:      strconv.Itoa(v.Rid),
				Type:        params.Type,
				RoomName:    v.RoomName,
				OwnerName:   v.NickName,
				OwnerAvatar: v.Avatar,
				IsLive:      true,
				RoomPic:     v.RoomSrc,
			})
		}
		result.Total = res.NumFound

		return &result, nil
	default:
		return nil, errors.New(TypeErr)
	}
}
