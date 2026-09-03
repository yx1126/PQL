package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	clouddirve "pql/pkg/cloud_dirve"
	"pql/pkg/vo"
	"sync"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type DriveService struct {
	*ServiceContext
	baidu *clouddirve.BaiduDrive
}

func NewDriveService(sc *ServiceContext) *DriveService {
	return &DriveService{
		ServiceContext: sc,
		baidu:          clouddirve.NewBaidu(sc.Http, sc.Auth),
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

func (s *DriveService) StartBaiduAuth() (*clouddirve.BaiduDeviceRes, error) {
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

	var result clouddirve.BaiduDeviceRes
	if err != json.Unmarshal(resp.Bytes(), &result) {
		return nil, err
	}

	return &result, nil
}

func (s *DriveService) GetAuthList() ([]vo.AuthVo, error) {
	res := s.Auth.GetAuthListist()
	var errs []error
	var wg sync.WaitGroup
	for i, v := range res {
		wg.Go(func() {
			if v.IsAuth() && v.Type == "baidu" {
				bd, err := s.baidu.GetInfo()
				if err != nil {
					errs = append(errs, err)
				} else {
					res[i].Avatar = bd.AvatarUrl
					res[i].DriveId = bd.Uk
					res[i].Username = bd.BaiduName
					res[i].Nickname = bd.NetdiskName
					res[i].VipType = bd.VipType
				}
			}
		})
	}
	wg.Wait()
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return res, nil
}

// 保存授权信息
func (s *DriveService) saveBaiduAuth(res *clouddirve.BaiduTokenRes) error {
	return s.Auth.SaveAuth(vo.SaveAuthVo{
		BaseAuth: vo.BaseAuth{
			Type:         "baidu",
			Token:        res.Token,
			ExpiresIn:    res.ExpiresIn,
			ExpiresTime:  time.Now().Add(time.Second * time.Duration(res.ExpiresIn)).Format(time.DateTime),
			RefreshToken: res.RefreshToken,
			Scope:        res.Scope,
		},
	})
}

// 获取百度token授权
func (s *DriveService) GetBaiduToken(code string) (*clouddirve.BaiduTokenRes, error) {
	res, err := s.baidu.GetAuthToken(code)
	if err != nil {
		return nil, err
	}
	if err := s.saveBaiduAuth(res); err != nil {
		return nil, err
	}
	return res, nil
}

// 刷新百度授权
func (s *DriveService) RefreshBaiduToken(refreshToken string) (*clouddirve.BaiduTokenRes, error) {
	res, err := s.baidu.RefreshAuthToken(refreshToken)
	if err != nil {
		return nil, err
	}
	if err := s.saveBaiduAuth(res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DriveService) UnBindBaidu(typee string) error {
	return s.Auth.DeleteAuth(typee)
}
