package service

import (
	"context"
	"pql/pkg/constant"
	"pql/pkg/utils/types"
	"sync"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

type WindowService struct {
	*ServiceContext
	mu     sync.RWMutex
	winMap map[string]*application.WebviewWindow
}

func NewWindowService(sc *ServiceContext) *WindowService {
	return &WindowService{
		ServiceContext: sc,
		winMap:         make(map[string]*application.WebviewWindow),
	}
}

func RegisterWindowStatusEvent(app *application.App, window *application.WebviewWindow) {
	name := window.Name()
	window.OnWindowEvent(events.Common.WindowMaximise, func(event *application.WindowEvent) {
		app.Event.Emit(constant.WindowMaximise, types.WindowState{
			Type: "maximise",
			Name: name,
		})
	})
	window.OnWindowEvent(events.Common.WindowUnMaximise, func(event *application.WindowEvent) {
		app.Event.Emit(constant.WindowUnMaximise, types.WindowState{
			Type: "unMaximise",
			Name: name,
		})
	})
	window.OnWindowEvent(events.Common.WindowRestore, func(event *application.WindowEvent) {
		app.Event.Emit(constant.WindowRestore, types.WindowState{
			Type: "restore",
			Name: name,
		})
	})
}

// Initialisation
func (ws *WindowService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return nil
}

func (ws *WindowService) ServiceShutdown() error {
	return nil
}

func (ws *WindowService) OpenNewWindow(options types.WindowOptions) {
	if options.Name == "" {
		return
	}
	ws.mu.RLock()
	winItem, ok := ws.winMap[options.Name]
	ws.mu.RUnlock()
	if ok {
		winItem.SetURL(options.Path).SetTitle(options.Title).Show().Focus()
		return
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()
	window := ws.App.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: options.Title,
		Name:  options.Name,
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: false,
		},
		Frameless:        true,
		URL:              options.Path,
		BackgroundColour: application.NewRGBA(15, 15, 22, 1),
		Width:            1200,
		Height:           800,
		MinWidth:         1200,
		MinHeight:        800,
	})
	ws.winMap[options.Name] = window
	window.OnWindowEvent(events.Common.WindowDidResize, func(event *application.WindowEvent) {
		window.SetMinSize(1200, 800)
	})
	window.OnWindowEvent(events.Common.WindowClosing, func(event *application.WindowEvent) {
		ws.mu.Lock()
		delete(ws.winMap, options.Name)
		ws.mu.Unlock()
	})
	RegisterWindowStatusEvent(ws.App, window)
}

func (ws *WindowService) Close(name string) {
	ws.mu.RLock()
	win, ok := ws.winMap[name]
	ws.mu.RUnlock()
	if ok {
		win.Close()
	}
}

func (ws *WindowService) GetAllWindows() []string {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	result := make([]string, 0, len(ws.winMap))
	for key := range ws.winMap {
		result = append(result, key)
	}
	return result
}

func (ws *WindowService) CloseAll() {
	ws.mu.RLock()
	windows := make([]*application.WebviewWindow, 0, len(ws.winMap))
	for _, win := range ws.winMap {
		windows = append(windows, win)
	}
	ws.mu.RUnlock()
	for _, win := range windows {
		win.Close()
	}
}
