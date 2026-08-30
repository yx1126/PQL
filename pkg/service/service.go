package service

import (
	"pql/DB"
	"pql/DB/model"
	ds "pql/DB/service"
	"pql/pkg/request"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type ServiceContext struct {
	App    *application.App
	Window *application.WebviewWindow
	Http   *request.Http
	DB     *DB.Sqlite

	// db
	Game   *ds.GameService
	Live   *ds.LiveService
	Parser *ds.ParserService
	Set    *ds.SettingService
	Menu   *ds.MenuService
	Store  *ds.StoreService
}

func (sc *ServiceContext) open() error {
	if err := sc.DB.Open("PQL.db"); err != nil {
		return err
	}
	return nil
}

func (sc *ServiceContext) autoMigrate() {
	sc.DB.AutoMigrate(
		&model.Game{},
		&model.Live{},
		&model.Parser{},
		&model.Setting{},
		&model.Menu{},
		&model.Store{},
	)
}

func (sc *ServiceContext) Close() error {
	if err := sc.DB.Close(); err != nil {
		return err
	}
	if err := sc.Http.Close(); err != nil {
		return err
	}
	return nil
}

func New(app *application.App, window *application.WebviewWindow) (*ServiceContext, error) {

	// http
	http := request.New()
	http.SetTimeout(2 * time.Minute)

	// db
	ns := DB.New(app)

	// context
	ctx := app.Context()

	// ServiceContext
	sc := &ServiceContext{
		App:    app,
		Window: window,
		Http:   http,
		DB:     ns,

		// services
		Game:   ds.NewGameService(ns, ctx),
		Live:   ds.NewLiveService(ns, ctx),
		Parser: ds.NewParserService(ns, ctx),
		Set:    ds.NewSettingService(ns, ctx),
		Menu:   ds.NewMenuService(ns, ctx),
		Store:  ds.NewStoreService(ns, ctx),
	}

	if err := sc.open(); err != nil {
		if err := sc.Http.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}

	// 初始化
	sc.autoMigrate()

	app.RegisterService(application.NewServiceWithOptions(
		NewGameService(sc),
		application.ServiceOptions{
			Name: "Game",
		},
	))
	app.RegisterService(application.NewServiceWithOptions(
		NewWindowService(sc),
		application.ServiceOptions{
			Name: "Window",
		},
	))
	app.RegisterService(application.NewServiceWithOptions(
		NewSetService(sc),
		application.ServiceOptions{
			Name: "Set",
		},
	))
	app.RegisterService(application.NewServiceWithOptions(
		NewAppService(sc),
		application.ServiceOptions{
			Name: "App",
		},
	))
	app.RegisterService(application.NewServiceWithOptions(
		NewParserService(sc),
		application.ServiceOptions{
			Name: "Parser",
		},
	))
	app.RegisterService(application.NewServiceWithOptions(
		NewFileServer(sc),
		application.ServiceOptions{
			Name:  "File",
			Route: "/local",
		},
	))
	app.RegisterService(application.NewServiceWithOptions(
		NewLiveService(sc),
		application.ServiceOptions{
			Name: "Live",
		},
	))

	return sc, nil
}
