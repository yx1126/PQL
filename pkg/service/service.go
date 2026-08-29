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
	app  *application.App
	http *request.Http
	db   *DB.Sqlite

	// db
	game   *ds.GameService
	live   *ds.LiveService
	parser *ds.ParserService
	set    *ds.SettingService
	menu   *ds.MenuService
	store  *ds.StoreService
}

func (sc *ServiceContext) open() error {
	if err := sc.db.Open("PQL.db"); err != nil {
		return err
	}
	return nil
}

func (sc *ServiceContext) autoMigrate() {
	sc.db.AutoMigrate(
		&model.Game{},
		&model.Live{},
		&model.Parser{},
		&model.Setting{},
		&model.Menu{},
		&model.Store{},
	)
}

func (sc *ServiceContext) Close() error {
	if err := sc.db.Close(); err != nil {
		return err
	}
	if err := sc.http.Close(); err != nil {
		return err
	}
	return nil
}

func New(app *application.App) (*ServiceContext, error) {

	// http
	http := request.New()
	http.SetTimeout(2 * time.Minute)

	// db
	ns := DB.New(app)

	// context
	ctx := app.Context()

	// ServiceContext
	sc := &ServiceContext{
		app:  app,
		http: http,
		db:   ns,

		// services
		game:   ds.NewGameService(ns, ctx),
		live:   ds.NewLiveService(ns, ctx),
		parser: ds.NewParserService(ns, ctx),
		set:    ds.NewSettingService(ns, ctx),
		menu:   ds.NewMenuService(ns, ctx),
		store:  ds.NewStoreService(ns, ctx),
	}

	if err := sc.open(); err != nil {
		if err := sc.http.Close(); err != nil {
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
