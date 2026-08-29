package main

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"pql/pkg/constant"
	"pql/pkg/logger"
	"pql/pkg/service"
	"pql/pkg/utils/instance"
	"pql/pkg/utils/tool"
	"pql/pkg/utils/types"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"github.com/wailsapp/wails/v3/pkg/services/dock"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	// Register a custom event whose associated data type is string.
	// This is not required, but the binding generator will pick up registered events
	// and provide a strongly typed JS/TS API for them.

	// 窗口状态事件
	application.RegisterEvent[types.WindowState](constant.WindowMaximise)
	application.RegisterEvent[types.WindowState](constant.WindowUnMaximise)
	application.RegisterEvent[types.WindowState](constant.WindowRestore)
	// 主题切换事件
	application.RegisterEvent[types.WindowTheme](constant.AppTheme)
	// 页面跳转事件
	application.RegisterEvent[types.PageChange](constant.PageChange)
}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {

	var mainWindow *application.WebviewWindow

	insts := instance.New()

	var zlog *logger.Logger

	app := application.New(application.Options{
		Name:        "PQL",
		Description: "Player Quick Launcher",
		Services: []application.Service{
			application.NewService(dock.New()),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Windows: application.WindowsOptions{},
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "com.yx1126.pql",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				if mainWindow != nil {
					mainWindow.Show().Focus()
				}
			},
		},
		RawMessageHandler: func(window application.Window, message string, originInfo *application.OriginInfo) {
			if zlog != nil {
				msg := fmt.Sprintf("Window [%s] with message: %s, origin %s, topOrigin %s, isMainFrame %t", window.Name(), message, originInfo.Origin, originInfo.TopOrigin, originInfo.IsMainFrame)
				zlog.Error(msg, errors.New(""))
			}
		},
	})

	if !app.Env.Info().Debug {
		zlog = logger.New(logger.ErrorPath)
		insts.Add(zlog)
	}

	// 服务注册
	sc, err := service.New(app)
	if err != nil {
		insts.Close()
		log.Fatal(err)
	}
	insts.Add(sc)

	mainWindow = app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "PQL",
		Name:  "main",
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: false,
		},
		Frameless:        true,
		BackgroundType:   application.BackgroundTypeTransparent,
		BackgroundColour: application.NewRGB(10, 11, 20),
		URL:              "/",
		Width:            1200,
		Height:           800,
		MinWidth:         1200,
		MinHeight:        800,
		DevToolsEnabled:  true,
	})

	// 托盘注册
	trayApp := NewSystray(app, mainWindow)
	insts.Add(trayApp)

	// 监听 app 退出
	app.OnShutdown(func() {
		insts.Close()
	})
	// 监听主题
	app.Event.OnApplicationEvent(events.Common.ThemeChanged, func(event *application.ApplicationEvent) {
		app.Event.Emit(constant.AppTheme, types.WindowTheme{
			Type:  "system",
			Theme: tool.Flag(app.Env.IsDarkMode(), 0, 1),
		})
	})

	// 注册窗口事件
	service.RegisterWindowStatusEvent(app, mainWindow)
	// Run the application. This blocks until the application has been exited.
	// If an error occurred while running the application, log it and exit.
	if err := app.Run(); err != nil {
		insts.Close()
		log.Fatal(err)
	}
}
