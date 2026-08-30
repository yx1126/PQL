package main

import (
	_ "embed"
	"pql/pkg/constant"
	"pql/pkg/service"
	"pql/pkg/utils/types"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type trayApp struct {
	*service.ServiceContext
	systray *application.SystemTray
}

//go:embed build/appicon.png
var icon []byte

func NewSystray(sc *service.ServiceContext) *trayApp {
	t := &trayApp{
		ServiceContext: sc,
	}
	t.setup()
	t.createMenus()
	return t
}

func (t *trayApp) setup() {
	systray := t.App.SystemTray.New()

	systray.SetIcon(icon)
	systray.SetTooltip("PQL")
	systray.SetLabel("PQL")

	systray.OnClick(func() {
		if t.Window.IsMinimised() {
			t.Window.Restore()
		}
		t.Window.Show().Focus()
	})

	t.systray = systray
}

func (t *trayApp) setTheme(theme int) {
	t.App.Event.Emit(constant.AppTheme, types.WindowTheme{
		Type:  "*",
		Theme: theme,
	})
}

func (t *trayApp) createMenus() {
	menu := t.App.NewMenu()

	// theme
	theme := menu.AddSubmenu("主题")
	theme.Add("深色模式").OnClick(func(ctx *application.Context) {
		t.setTheme(0)
	})
	theme.Add("浅色模式").OnClick(func(ctx *application.Context) {
		t.setTheme(1)
	})
	theme.Add("跟随系统").OnClick(func(ctx *application.Context) {
		t.setTheme(2)
	})

	// 重置位置
	menu.Add("重置").OnClick(func(ctx *application.Context) {
		t.Window.Center()
		t.Window.Restore()
		t.Window.Show().Focus()
	})

	// setting
	menu.Add("设置").OnClick(func(ctx *application.Context) {
		t.App.Event.Emit(constant.PageChange, types.PageChange{
			Type: "push",
			Path: "/sub/setting",
		})
		if t.Window.IsMinimised() {
			t.Window.Restore()
		}
		t.Window.Show().Focus()
	})

	// split
	menu.AddSeparator()

	// quit
	menu.Add("退出").OnClick(func(ctx *application.Context) {
		t.App.Quit()
	})
	t.systray.SetMenu(menu)
}

func (t *trayApp) Close() error {
	if t.systray != nil {
		t.systray.Destroy()
	}
	return nil
}
