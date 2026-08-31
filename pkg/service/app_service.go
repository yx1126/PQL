package service

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
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

func (s *AppService) StartBaiduAuth() error {

	params := url.Values{}
	params.Add("response_type", "token")
	params.Add("client_id", "iV7sfG52vgnNTjPceUt2xCQNdfum6gJm")
	params.Add("redirect_uri", "http://127.0.0.1:58080/baidu")
	params.Add("scope", "basic,netdisk")

	if err := s.App.Browser.OpenURL("https://openapi.baidu.com/oauth/2.0/authorize?" + params.Encode()); err != nil {
		return err
	}
	// result := make(chan string, 1)
	// mux := http.NewServeMux()
	// server := &http.Server{Addr: "127.0.0.1:58080", Handler: mux}
	// mux.HandleFunc("/baidu", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodGet {
	// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// 		return
	// 	}
	// 	code := r.URL.Query().Get("code")

	// 	result <- code
	// 	go server.Shutdown(context.Background())
	// })
	// errCh := make(chan error, 1)
	// go func() {
	// 	err := server.ListenAndServe()

	// 	if err != nil && err != http.ErrServerClosed {
	// 		errCh <- err
	// 	}
	// }()
	// select {
	// case code := <-result:
	// 	return code, nil

	// case err := <-errCh:
	// 	return "", err

	// case <-time.After(5 * time.Minute):
	// 	_ = server.Shutdown(context.Background())
	// 	return "", fmt.Errorf("等待授权超时")
	// }
	return nil
}
