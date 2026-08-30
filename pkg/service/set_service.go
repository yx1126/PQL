package service

import (
	"context"
	"pql/pkg/vo"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type SetService struct {
	*ServiceContext
}

func NewSetService(sc *ServiceContext) *SetService {
	return &SetService{
		ServiceContext: sc,
	}
}

// Initialisation
func (s *SetService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	// setting init
	if err := s.Set.Init(); err != nil {
		return err
	}
	// menu init
	if err := s.Menu.Init(); err != nil {
		return err
	}
	// store init
	if err := s.Store.Init(); err != nil {
		return err
	}
	return nil
}

func (gs *SetService) ServiceShutdown() error {
	return nil
}

// setting
func (s *SetService) GetConfig() vo.SettingVo {
	return s.Set.GetSetting()
}

func (s *SetService) UpdateConfig(config vo.UpdateSettingVo) error {
	return s.Set.UpdateSetting(config)
}

// menu
func (s *SetService) GetMenuList() []vo.MenuVo {
	return s.Menu.GetMenuList()
}

func (s *SetService) UpdateMenu(menu vo.UpdateMenuVo) error {
	return s.UpdateMenu(menu)
}
