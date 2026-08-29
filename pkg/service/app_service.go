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
	// if a.app.Env.Info().Debug {
	// 	a.resty.SetDebug(true)
	// }
	return nil
}

func (a *AppService) ServiceShutdown() error {

	return nil
}

func (a *AppService) AutoStartStatus() application.AutostartStatus {
	st, err := a.app.Autostart.Status()
	if err != nil {
		return application.AutostartStatus{
			Enabled: false,
		}
	}
	return st
}

func (a *AppService) StartOnWindow() error {
	return a.app.Autostart.Enable()
}

func (a *AppService) DisableOnWindow() error {
	return a.app.Autostart.Disable()
}

func (a *AppService) GetDarkMode() bool {
	return a.app.Env.IsDarkMode()
}

func (a *AppService) Request(config request.HttpRequestConfig) any {

	res, err := a.http.Request(config)

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
