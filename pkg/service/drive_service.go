package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"pql/pkg/utils/types"
	"pql/pkg/vo"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type DriveService struct {
	*ServiceContext
}

func NewDriveService(sc *ServiceContext) *DriveService {
	return &DriveService{
		ServiceContext: sc,
	}
}

// Initialisation
func (a *DriveService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	if err := a.Auth.Init(); err != nil {
		return err
	}
	return nil
}

func (a *DriveService) ServiceShutdown() error {

	return nil
}

func (s *DriveService) StartBaiduAuth() (*types.BaiduDeviceRes, error) {
	r := s.Http.R()

	params := url.Values{}
	params.Add("response_type", "device_code")
	params.Add("client_id", "iV7sfG52vgnNTjPceUt2xCQNdfum6gJm")
	params.Add("scope", "basic,netdisk")
	params.Add("response_type", "device_code")

	r.SetQueryParamsFromValues(params)

	r.SetHeader("User-Agent", "pan.baidu.com")

	resp, err := r.Get("https://openapi.baidu.com/oauth/2.0/device/code")
	if err != nil {
		return nil, err
	}

	var result types.BaiduDeviceRes
	if err != json.Unmarshal(resp.Bytes(), &result) {
		return nil, err
	}

	return &result, nil
}

func (s *DriveService) GetAuthList() []vo.AuthVo {
	return s.Auth.GetAuthListist()
}

func (s *DriveService) GetAuth(typee string) (*vo.AuthVo, error) {
	return s.Auth.GetAuth(typee)
}

func (s *DriveService) getBaiduTokens(params url.Values) (*types.BaiduTokenRes, error) {
	r := s.Http.R()

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
		var baiduErr types.BaiduAuthError
		if err != json.Unmarshal(resp.Bytes(), &baiduErr) {
			return nil, err
		}
		return nil, errors.New(baiduErr.ErrorDescription)
	}

	var result types.BaiduTokenRes
	if err != json.Unmarshal(resp.Bytes(), &result) {
		return nil, err
	}
	time.Now().Format(time.DateTime)
	if err := s.Auth.SaveAuth(vo.SaveAuthVo{
		BaseAuth: vo.BaseAuth{
			Type:         "baidu",
			Token:        result.AccessToken,
			ExpiresIn:    result.ExpiresIn,
			ExpiresTime:  time.Now().Add(time.Second * time.Duration(result.ExpiresIn)).Format(time.DateTime),
			RefreshToken: result.RefreshToken,
			Scope:        result.Scope,
		},
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *DriveService) GetBaiduToken(code string) (*types.BaiduTokenRes, error) {
	params := url.Values{}
	params.Add("grant_type", "device_token")
	params.Add("code", code)
	return s.getBaiduTokens(params)
}

func (s *DriveService) RefreshBaiduToken(typee, refreshToken string) (*types.BaiduTokenRes, error) {
	params := url.Values{}
	params.Add("grant_type", "refresh_token")
	params.Add("refresh_token", refreshToken)
	return s.getBaiduTokens(params)
}

func (s *DriveService) UnBindBaidu(typee string) error {
	return s.Auth.DeleteAuth(typee)
}
