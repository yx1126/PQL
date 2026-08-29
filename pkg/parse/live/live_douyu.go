package live

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"pql/pkg/parse/sign"
	"pql/pkg/request"
	"pql/pkg/utils/tool"
	"strconv"
	"time"

	"github.com/dop251/goja"
)

const (
	did = "10000000000000000000000000001501"
)

var teaKey = [4]uint32{
	0x05234bd7,
	0x57362bb4,
	0x00720000,
	0x7f140c50,
}

type PlayInfo struct {
	RtmpCdn    string `json:"rtmp_cdn"`
	RtmpLive   string `json:"rtmp_live"`
	RtmpUrl    string `json:"rtmp_url"`
	Rate       int    `json:"rate"`
	Multirates []struct {
		Name string `json:"name"`
		Rate int    `json:"rate"`
	} `json:"multirates"`
	CdnsWithName []CdnsWithName `json:"cdnsWithName"`
}

type DouyuResponse[T any] struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
	Data  T      `json:"data"`
}

type Douyu struct {
	userAgent string
	http      *request.Http
}

type DouyuRoomInfo struct {
	Room struct {
		RoomId      int    `json:"room_id"`
		RoomName    string `json:"room_name"`
		ShowStatus  int    `json:"show_status"`
		VideoLoop   int    `json:"videoLoop"`
		OwnerName   string `json:"owner_name"`
		OwnerAvatar string `json:"owner_avatar"`
		RoomPic     string `json:"room_pic"`
		CoverSrc    string `json:"coverSrc"`
		ShowDetails string `json:"show_details"`
	} `json:"room"`
}

type DouyuSearchInfo struct {
	PageSize   int `json:"pageSize"`
	RelateShow []struct {
		Rid      int    `json:"rid"`
		NickName string `json:"nickName"`
		Avatar   string `json:"avatar"`
		CateName string `json:"cateName"`
		RoomName string `json:"roomName"`
		RoomSrc  string `json:"roomSrc"`
	} `json:"relateShow"`
	Total int `json:"total"`
}

func NewDouyu(http *request.Http) *Douyu {
	return &Douyu{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/151.0.0.0 Safari/537.36",
		http:      http,
	}
}

func (d *Douyu) GetHeaders(roomId string) map[string]string {
	return map[string]string{
		"accept":     "application/json, text/plain, */*",
		"referer":    fmt.Sprintf("https://www.douyu.com/%v", roomId),
		"user-agent": d.userAgent,
	}
}

func (d *Douyu) GetRoomWebUri(roomId string) string {
	return "https://www.douyu.com/" + roomId
}

func (d *Douyu) GetInfo(roomId string) (*RoomInfo, error) {
	r := d.http.R()

	r.SetHeaders(d.GetHeaders(roomId))

	resp, err := r.Execute(http.MethodGet, fmt.Sprintf("https://www.douyu.com/betard/%v", roomId))
	if err != nil {
		return nil, err
	}

	var dyInfo *DouyuRoomInfo
	if err := json.Unmarshal(resp.Bytes(), &dyInfo); err == nil {
		r := dyInfo.Room
		return &RoomInfo{
			RoomId:      strconv.Itoa(r.RoomId),
			RoomName:    r.RoomName,
			OwnerName:   r.OwnerName,
			IsLive:      r.ShowStatus == 1,
			IsReplay:    r.VideoLoop == 1,
			OwnerAvatar: r.OwnerAvatar,
			RoomPic:     tool.Flag(r.CoverSrc != "", r.CoverSrc, r.RoomPic),
			Description: r.ShowDetails,
		}, nil
	}

	return nil, errors.New(RoomErr)
}

func (d *Douyu) GetPlayUrl(roomId string, extra map[string]string) (*LiveInfo, error) {
	r := d.http.R()

	script, err := d.getSignScript(roomId)
	if err != nil {
		return nil, err
	}

	signStr, err := d.createSign(script, roomId)
	if err != nil {
		return nil, err
	}

	sign, err := url.ParseQuery(signStr)
	if err != nil {
		return nil, err
	}

	r.SetFormDataFromValues(sign)
	r.SetFormData(extra)
	r.SetHeaders(d.GetHeaders(roomId))

	resp, err := r.Execute(http.MethodPost, fmt.Sprintf("https://www.douyu.com/lapi/live/getH5Play/%v", roomId))
	if err != nil {
		return nil, err
	}
	res := DouyuResponse[PlayInfo]{}

	if err := json.Unmarshal(resp.Bytes(), &res); err == nil {
		rd := res.Data
		multirates := make([]Multirate, len(rd.Multirates))
		for i, v := range rd.Multirates {
			multirates[i] = Multirate{
				Name: v.Name,
				Rate: strconv.Itoa(v.Rate),
			}
		}
		return &LiveInfo{
			RoomId:       roomId,
			VideoType:    "flv",
			Url:          rd.RtmpUrl + "/" + rd.RtmpLive,
			RtmpCdn:      rd.RtmpCdn,
			Rate:         strconv.Itoa(rd.Rate),
			Multirates:   multirates,
			CdnsWithName: rd.CdnsWithName,
		}, nil
	}

	errRes := DouyuResponse[string]{}
	if err := json.Unmarshal(resp.Bytes(), &errRes); err == nil {
		return nil, errors.New(errRes.Msg)
	}
	return nil, errors.New("未知错误！")
}

func (d *Douyu) Search(params LiveParams) (*DouyuSearchInfo, error) {
	r := d.http.R()

	r.SetHeaders(map[string]string{
		"user-agent": d.userAgent,
		"referer":    "https://www.douyu.com/search?&kw=" + params.Keyword,
		"Cookie":     fmt.Sprintf("dy_did=%v; acf_did=%v;", did, did),
	})

	r.SetQueryParams(map[string]string{
		"page":     params.Page,
		"pageSize": "20",
		"kw":       params.Keyword,
	})

	resp, err := r.Execute(http.MethodGet, "https://www.douyu.com/japi/search/api/searchShow")
	if err != nil {
		return nil, err
	}

	var result DouyuResponse[DouyuSearchInfo]

	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, err
	}

	if result.Error == 8 {
		return nil, errors.New(result.Msg)
	}

	return &result.Data, nil
}

func (d *Douyu) getSignScript(roomId string) (string, error) {
	r := d.http.R()

	r.SetHeaders(d.GetHeaders(roomId))

	resp, err := r.Execute(http.MethodGet, fmt.Sprintf("https://www.douyu.com/swf_api/homeH5Enc?rids=%v", roomId))
	if err != nil {
		return "", errors.New("斗鱼没有返回 homeH5Enc 签名脚本")
	}

	res := &DouyuResponse[map[string]string]{}
	if err := json.Unmarshal(resp.Bytes(), &res); err != nil {
		return "", err
	}

	str, ok := res.Data[fmt.Sprintf("room%v", roomId)]
	if !ok {
		return "", fmt.Errorf("斗鱼没有返回 homeH5Enc 签名脚本")
	}

	return str, nil
}

func (d *Douyu) createSign(script string, roomId string) (string, error) {
	vm := goja.New()

	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli()/1000)

	sc := fmt.Sprintf(`(() => { %v; %v; return ub98484234("%v","%v","%v") })()`,
		sign.Cryptojs,
		script,
		roomId,
		did,
		timestamp,
	)

	val, err := vm.RunString(sc)
	if err != nil {
		return "", err
	}

	return val.String(), err
}
