package service

import (
	"context"
	"pql/pkg/vo"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type GameService struct {
	*ServiceContext
	ctx context.Context
}

func NewGameService(sc *ServiceContext) *GameService {
	return &GameService{
		ServiceContext: sc,
	}
}

// Initialisation
func (g *GameService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	g.ctx = ctx
	if err := g.Game.Init(); err != nil {
		return err
	}
	return nil
}

func (gs *GameService) ServiceShutdown() error {
	return nil
}

func (g *GameService) GetGameList() []vo.GameVo {
	return g.Game.GetGameList()
}

func (g *GameService) UpdateGame(game vo.UpdateGameVo) error {
	return g.Game.UpdateGame(game)
}
