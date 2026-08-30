package live

import (
	"encoding/json"
	"errors"
	"net/http"
	"pql/pkg/request"
	"pql/pkg/utils/tool"
	"strconv"
	"strings"
)

const (
	// 匹配19位数字 room_id
	Ttwid = "1%7CWNA0hGL9x5b8iy2MWgm_s4NLI6rCSwNWSVjcaQnlpVQ%7C1787887833%7C862fcc45187368aa0cd47cef457174452263ee807d3e3663f110850ff124b06c"
)

type (
	Douyin struct {
		userAgent string
		http      *request.Http
	}

	DouyinResponse[T any] struct {
		StatusCode int    `json:"status_code"`
		StatusMsg  string `json:"status_msg"`
		Data       T      `json:"data"`
	}

	DouyinUserResponse[T any] struct {
		StatusCode int     `json:"status_code"`
		StatusMsg  *string `json:"status_msg"`
		User       T       `json:"user"`
	}

	DouyinQualities struct {
		Name   string `json:"name"`
		SdkKey string `json:"sdk_key"`
	}

	DouyinLine struct {
		Flv string `json:"flv"`
		Hls string `json:"hls"`
	}

	DouyinStream struct {
		Common struct {
			Lines map[string]string `json:"lines"`
		} `json:"common"`
		Data map[string](map[string]DouyinLine) `json:"data"`
	}

	DouyinRoomInfo struct {
		Room struct {
			Status int `json:"status"`

			RoomId   string `json:"id_str"`
			RoomName string `json:"title"`

			Cover struct {
				UrlList []string `json:"url_list"`
			} `json:"cover"`

			Stream struct {
				ResolutionName    map[string]string `json:"resolution_name"`
				DefaultResolution string            `json:"default_resolution"`

				RtmpPullUrl string            `json:"rtmp_pull_url"`
				FlvPullUrl  map[string]string `json:"flv_pull_url"`

				HLSPullUrl string `json:"hls_pull_url"`

				HlsPullUrlMap map[string]string `json:"hls_pull_url_map"`

				LiveCoreSdkData struct {
					PullData struct {
						StreamData string `json:"stream_data"`
						Options    struct {
							DefaultQuality DouyinQualities   `json:"default_quality"`
							Qualities      []DouyinQualities `json:"qualities"`
						} `json:"options"`
					} `json:"pull_data"`
				} `json:"live_core_sdk_data"`
			} `json:"stream_url"`

			Owner struct {
				Nickname    string `json:"nickname"`
				Signature   string `json:"signature"`
				AvatarThumb struct {
					UrlList []string `json:"url_list"`
					Uri     string   `json:"uri"`
				} `json:"avatar_thumb"`
				FollowInfo struct {
					// 粉丝数
					FollowerCount    int    `json:"follower_count"`
					FollowerCountStr string `json:"follower_count_str"`
				} `json:"follow_info"`
			} `json:"owner"`
		} `json:"room"`
		Message string `json:"message"`
	}

	DouyinUserInfo struct {
		RoomId      string `json:"room_id_str"`
		Nickname    string `json:"nickname"`
		Signature   string `json:"signature"`
		AvatarThumb struct {
			UrlList []string `json:"url_list"`
		} `json:"avatar_thumb"`

		Room struct {
			Status int `json:"status"`

			RoomId      string `json:"id_str"`
			RoomName    string `json:"title"`
			ShowStatus  int    `json:"show_status"`
			VideoLoop   int    `json:"videoLoop"`
			OwnerName   string `json:"owner_name"`
			OwnerAvatar string `json:"owner_avatar"`
			RoomPic     string `json:"room_pic"`
			CoverSrc    string `json:"coverSrc"`
			ShowDetails string `json:"show_details"`

			Stream struct {
				ResolutionName    map[string]string `json:"resolution_name"`
				DefaultResolution string            `json:"default_resolution"`

				RtmpPullUrl string            `json:"rtmp_pull_url"`
				FlvPullUrl  map[string]string `json:"flv_pull_url"`

				HLSPullUrl string `json:"hls_pull_url"`

				HlsPullUrlMap map[string]string `json:"hls_pull_url_map"`
			} `json:"stream_url"`

			Owner struct {
				Nickname    string `json:"nickname"`
				Signature   string `json:"signature"`
				AvatarThumb struct {
					UrlList []string `json:"url_list"`
					Uri     string   `json:"uri"`
				} `json:"avatar_thumb"`
				FollowInfo struct {
					// 粉丝数
					FollowerCount    int    `json:"follower_count"`
					FollowerCountStr string `json:"follower_count_str"`
				} `json:"follow_info"`
			} `json:"owner"`
		} `json:"room"`
		Message string `json:"message"`
	}
)

func NewDouyin(http *request.Http) *Douyin {
	return &Douyin{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/151.0.0.0 Safari/537.36",
		http:      http,
	}
}

func (d *Douyin) GetRoomWebUri(roomId string) string {
	return "https://www.douyin.com/user/" + roomId
}

func (d *Douyin) GetInfo(roomId, ttwid string) (*RoomInfo, error) {
	r := d.http.R()

	r.SetHeaders(map[string]string{
		"referer":    "https://www.douyin.com/user/" + roomId,
		"cookie":     "ttwid=" + ttwid,
		"User-Agent": d.userAgent,
	})

	r.SetQueryParam("sec_user_id", roomId)

	resp, err := r.Execute(http.MethodGet, "https://www.douyin.com/aweme/v1/web/user/profile/other")
	if err != nil {
		return nil, err
	}
	var result DouyinUserResponse[DouyinUserInfo]

	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, err
	}

	if result.StatusCode != 0 {
		return nil, errors.New(tool.Flag(result.StatusMsg != nil && *result.StatusMsg != "", *result.StatusMsg, UnknowErr))
	}

	u := result.User

	var (
		roomName string = u.Nickname
		roomPic  string
	)

	if u.RoomId != "0" {
		r, err := d.getBaseInfo(u.RoomId)
		if err == nil {
			roomName = r.Data.Room.RoomName
			pics := r.Data.Room.Cover.UrlList
			if len(pics) > 0 {
				roomPic = pics[0]
			}
		}
	}

	avatar := tool.Flag(len(u.AvatarThumb.UrlList) > 0, u.AvatarThumb.UrlList[0], "")

	return &RoomInfo{
		RoomId:      u.RoomId,
		Type:        "3",
		RoomName:    roomName,
		RoomPic:     roomPic,
		OwnerName:   u.Nickname,
		IsLive:      u.RoomId != "0",
		OwnerAvatar: avatar,
		Description: u.Signature,
	}, nil
}

func (d *Douyin) GetPlayUrl(roomId string, extra map[string]string) (*LiveInfo, error) {
	cdn := extra["cdn"]
	rate := extra["rate"]

	res, err := d.getBaseInfo(roomId)
	if err != nil {
		return nil, err
	}

	rd := res.Data.Room

	var stream DouyinStream
	if err := json.Unmarshal([]byte(rd.Stream.LiveCoreSdkData.PullData.StreamData), &stream); err != nil {
		return nil, err
	}

	// 线路
	cnds := []CdnsWithName{}
	count := 0
	for k, _ := range stream.Common.Lines {
		count++
		cnds = append(cnds, CdnsWithName{
			Name: "线路" + strconv.Itoa(count),
			Cdn:  k + "-flv",
		})
	}

	// 设置默认值
	if strings.TrimSpace(cdn) == "" && len(cnds) > 0 {
		cdn = cnds[0].Cdn
	}

	if strings.TrimSpace(rate) == "-1" {
		rate = rd.Stream.LiveCoreSdkData.PullData.Options.DefaultQuality.SdkKey
	}

	// 清晰度
	multirates := []Multirate{}
	qualities := rd.Stream.LiveCoreSdkData.PullData.Options.Qualities
	for i := len(qualities) - 1; i >= 0; i-- {
		multirates = append(multirates, Multirate{
			Name: qualities[i].Name,
			Rate: qualities[i].SdkKey,
		})
	}

	lineMap, ok := stream.Data[rate]
	if !ok {
		return nil, errors.New(ParserErr)
	}

	before, _, found := strings.Cut(cdn, "-")
	if !found {
		return nil, errors.New(ParserErr)
	}

	lines, ok := lineMap[before]
	if !ok {
		return nil, errors.New(ParserErr)
	}

	return &LiveInfo{
		RoomId:       rd.RoomId,
		VideoType:    "flv",
		Url:          lines.Flv,
		RtmpCdn:      cdn,
		Rate:         rate,
		CdnsWithName: cnds,
		Multirates:   multirates,
	}, nil
}

func (d *Douyin) getBaseInfo(roomId string) (*DouyinResponse[DouyinRoomInfo], error) {
	r := d.http.R()

	r.SetQueryParams(map[string]string{
		"type_id": "0",
		"live_id": "1",
		"room_id": roomId,
		"app_id":  "6383",
	})

	resp, err := r.Execute(http.MethodGet, "https://webcast.amemv.com/webcast/room/reflow/info")
	if err != nil {
		return nil, err
	}

	var result DouyinResponse[DouyinRoomInfo]

	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, err
	}

	if result.StatusCode != 0 {
		msg := result.Data.Message
		return nil, errors.New(tool.Flag(msg != "", msg, UnknowErr))
	}
	return &result, nil
}
