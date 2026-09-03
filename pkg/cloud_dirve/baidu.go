package clouddirve

import (
	"encoding/json"
	"errors"
	"net/url"
	"pql/DB/service"
	"pql/pkg/request"
	"pql/pkg/vo"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

type BaiduDeviceRes struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationUrl string `json:"verification_url"`
	Qrcodeurl       string `json:"qrcode_url"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

type BaiduTokenRes struct {
	Token        string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	ExpiresTime  string `json:"expires_time"`
	Scope        string `json:"scope"`
}

type BaiduAuthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type BaiduBaseRes struct {
	Errno  int    `json:"errno"`
	Errmsg string `json:"errmsg"`
}

type BaiduUserInfo struct {
	BaiduBaseRes
	Uk          int    `json:"uk"`
	AvatarUrl   string `json:"avatar_url"`
	BaiduName   string `json:"baidu_name"`
	NetdiskName string `json:"netdisk_name"`
	VipType     int    `json:"vip_type"`
}

type BaiduDrive struct {
	http *request.Http
	auth *service.AuthService

	mu sync.RWMutex
	sf singleflight.Group
}

func NewBaidu(http *request.Http, auth *service.AuthService) *BaiduDrive {
	result := &BaiduDrive{
		http: http,
		auth: auth,
	}
	return result
}

func (bd *BaiduDrive) GetInfo() (*BaiduUserInfo, error) {
	token, err := bd.GetToken()
	if err != nil {
		return nil, err
	}
	r := bd.http.R()
	r.SetQueryParams(map[string]string{
		"method":       "uinfo",
		"access_token": token,
		"vip_version":  "v2",
	})
	resp, err := r.Get("https://pan.baidu.com/rest/2.0/xpan/nas")
	if err != nil {
		return nil, err
	}
	var user BaiduUserInfo
	if err := json.Unmarshal(resp.Bytes(), &user); err != nil {
		return nil, err
	}
	if user.Errno != 0 {
		return nil, errors.New(user.Errmsg)
	}
	return &user, nil
}

func (bd *BaiduDrive) GetToken() (string, error) {
	baidu, err := bd.auth.GetAuth("baidu")
	if err != nil {
		return "", err
	}
	if !isExpired(baidu.ExpiresIn, baidu.ExpiresTime) {
		return baidu.Token, nil
	}
	v, err, _ := bd.sf.Do("baidu-token-refresh", func() (any, error) {
		baidu, err := bd.auth.GetAuth("baidu")
		if err != nil {
			return "", err
		}
		result, err := bd.RefreshAuthToken(baidu.RefreshToken)
		if err != nil {
			return "", err
		}
		bd.mu.Lock()
		defer bd.mu.Unlock()
		if err := bd.auth.SaveAuth(vo.SaveAuthVo{
			BaseAuth: vo.BaseAuth{
				Type:         "baidu",
				Token:        result.Token,
				ExpiresIn:    result.ExpiresIn,
				ExpiresTime:  result.ExpiresTime,
				RefreshToken: result.RefreshToken,
				Scope:        result.Scope,
			},
		}); err != nil {
			return "", err
		}
		return result.Token, nil
	})
	if err != nil {
		return "", err
	}
	return v.(string), nil
}

// 轮询获取token
func (bd *BaiduDrive) GetAuthToken(code string) (*BaiduTokenRes, error) {
	params := url.Values{}
	params.Add("grant_type", "device_token")
	params.Add("code", code)
	return bd.tokenAuth(params)
}

// 刷新 token
func (bd *BaiduDrive) RefreshAuthToken(refreshToken string) (*BaiduTokenRes, error) {
	params := url.Values{}
	params.Add("grant_type", "refresh_token")
	params.Add("refresh_token", refreshToken)
	return bd.tokenAuth(params)
}

// 获取/刷新 token
func (bd *BaiduDrive) tokenAuth(params url.Values) (*BaiduTokenRes, error) {
	r := bd.http.R()

	r.SetQueryParamsFromValues(params)

	baseParams := url.Values{}
	baseParams.Add("client_id", "iV7sfG52vgnNTjPceUt2xCQNdfum6gJm")
	baseParams.Add("client_secret", "28Q3eRQjJwtbRrjBpvAIqeFaOJCylUXG")
	r.SetQueryParamsFromValues(baseParams)

	r.SetHeader("User-Agent", "pan.baidu.com")

	resp, err := r.Get("https://openapi.baidu.com/oauth/2.0/token")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		var baiduErr BaiduAuthError
		if err := json.Unmarshal(resp.Bytes(), &baiduErr); err != nil {
			return nil, err
		}
		return nil, errors.New(baiduErr.ErrorDescription)
	}

	var result BaiduTokenRes
	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, err
	}
	result.ExpiresTime = time.Now().Add(time.Second * time.Duration(result.ExpiresIn)).Format(time.DateTime)

	return &result, nil
}

func isExpired(expiresIn int, expiresTime string) bool {
	if expiresIn == 0 || strings.TrimSpace(expiresTime) == "" {
		return true
	}
	expireTime, err := time.Parse(time.DateTime, expiresTime)
	if err != nil {
		return true
	}
	return time.Now().After(expireTime.Add(-5 * time.Minute))
}
