package service

import (
	"context"
	_ "embed"
	"pql/DB"
	"pql/DB/model"
	"pql/DB/scopes"
	"pql/pkg/vo"
)

type GameService struct {
	db  *DB.Sqlite
	ctx context.Context
}

func NewGameService(db *DB.Sqlite, ctx context.Context) *GameService {
	return &GameService{
		db:  db,
		ctx: ctx,
	}
}

func (g *GameService) Init() error {
	var count int64
	g.db.WithContext(g.ctx).Model(&model.Game{}).Count(&count)
	if count <= 0 {
		games := []model.Game{
			{Name: "七日杀", Type: "7days"},
			{Name: "PUBG", Type: "pubg", Path: "/pubg", IsSupportOpenWindow: 1},
		}
		if err := g.db.WithContext(g.ctx).Model(&model.Game{}).Create(&games).Error; err != nil {
			return err
		}
	}
	return nil
}

func (g *GameService) GetGameList() []vo.GameVo {
	var gameList = make([]vo.GameVo, 0)
	g.db.WithContext(g.ctx).Model(&model.Game{}).Select("*").Find(&gameList)
	return gameList
}

func (g *GameService) UpdateGame(game vo.UpdateGameVo) error {
	result := g.db.WithContext(g.ctx).
		Model(&model.Game{}).
		Scopes(scopes.UpdateOmitScope()).
		Where("id = ?", game.Id).
		Updates(&model.Game{
			CustomCommand: game.CustomCommand,
			InstallFolder: game.InstallFolder,
			ModsFolder:    game.ModsFolder,
			SaveFolder:    game.SaveFolder,
			Path:          game.Path,
			IsFixed:       game.IsFixed,
			Hidden:        game.Hidden,
		})
	return result.Error
}
