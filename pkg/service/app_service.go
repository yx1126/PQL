package service

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"os/exec"
	"pql/pkg/request"
	"pql/pkg/utils/types"
	"pql/pkg/vo"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type AppService struct {
	*ServiceContext
}

func NewAppService(sc *ServiceContext) *AppService {
	return &AppService{
		ServiceContext: sc,
	}
}

// Initialisation
func (a *AppService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	// if a.App.Env.Info().Debug {
	// 	a.resty.SetDebug(true)
	// }
	return nil
}

func (a *AppService) ServiceShutdown() error {

	return nil
}

func (a *AppService) AutoStartStatus() application.AutostartStatus {
	st, err := a.App.Autostart.Status()
	if err != nil {
		return application.AutostartStatus{
			Enabled: false,
		}
	}
	return st
}

func (a *AppService) StartOnWindow() error {
	return a.App.Autostart.Enable()
}

func (a *AppService) DisableOnWindow() error {
	return a.App.Autostart.Disable()
}

func (a *AppService) GetDarkMode() bool {
	return a.App.Env.IsDarkMode()
}

func (a *AppService) Request(config request.HttpRequestConfig) any {

	res, err := a.Http.Request(config)

	if err != nil {
		return err
	}

	var response any

	if err := json.Unmarshal(res.Bytes(), &response); err != nil {
		return nil
	}
	return response
}

func (a *AppService) HasHEVCExtension() bool {
	if runtime.GOOS != "windows" {
		return false
	}

	cmd := exec.Command(
		"powershell.exe",
		"-NoProfile",
		"-NonInteractive",
		"-Command",
		`Get-AppxPackage | Where-Object {$_.Name -in 'Microsoft.HEVCVideoExtension','Microsoft.HEVCVideoExtensions'}`,
	)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return false
	}

	return strings.TrimSpace(out.String()) != ""
}

// store
func (s *AppService) GetStore(key string) string {
	return s.Store.GetStore(key)
}

func (s *AppService) SetStore(key, value string) error {
	return s.Store.SetStore(key, value)
}

func (s *AppService) RemoveStore(key string) error {
	return s.Store.DeleteStore(key)
}

func (s *AppService) ClearStore() error {
	return s.Store.ClearStore()
}

func (s *AppService) StartBaiduAuth() (*types.BaiduDeviceRes, error) {
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

func (s *AppService) GetAuthList() []vo.AuthVo {
	return s.Auth.GetAuthListist()
}

func (s *AppService) GetAuth(typee string) (*vo.AuthVo, error) {
	return s.Auth.GetAuth(typee)
}

func (s *AppService) getBaiduTokens(params url.Values) (*types.BaiduTokenRes, error) {
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
			ExpiresTime:  time.Now().Add(time.Duration(result.ExpiresIn)).Format(time.DateTime),
			RefreshToken: result.RefreshToken,
			Scope:        result.Scope,
		},
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AppService) GetBaiduToken(code string) (*types.BaiduTokenRes, error) {
	params := url.Values{}
	params.Add("grant_type", "device_token")
	params.Add("code", code)
	return s.getBaiduTokens(params)
}

func (s *AppService) RefreshBaiduToken(typee, refreshToken string) (*types.BaiduTokenRes, error) {
	params := url.Values{}
	params.Add("grant_type", "refresh_token")
	params.Add("refresh_token", refreshToken)
	return s.getBaiduTokens(params)
}
