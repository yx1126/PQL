package DB

import (
	"context"
	"pql/pkg/utils/file"
	tool "pql/pkg/utils/tool"

	"github.com/glebarez/sqlite"
	"github.com/wailsapp/wails/v3/pkg/application"
	"gorm.io/gorm"
)

type Sqlite struct {
	app *application.App
	DB  *gorm.DB
}

func New(app *application.App) *Sqlite {
	return &Sqlite{
		app: app,
	}
}

func (s *Sqlite) WithContext(ctx context.Context) *gorm.DB {
	return s.DB.WithContext(ctx)
}

func (s *Sqlite) AutoMigrate(dst ...interface{}) error {
	return s.DB.AutoMigrate(dst...)
}

func (s *Sqlite) Open(name string) error {
	dir := tool.Flag(s.app.Env.Info().Debug, "bin/.PQL/dbs", ".PQL/dbs")
	if err := file.ValidDir(dir); err != nil {
		return err
	}
	dbPath := file.Joinwd(dir, name)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}
	s.DB = db
	return nil
}

func (s *Sqlite) Close() error {
	sqlDB, err := s.DB.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}
	return nil
}
