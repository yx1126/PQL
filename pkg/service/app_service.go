package service

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"pql/pkg/request"
	"runtime"
	"strings"
	"syscall"

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
