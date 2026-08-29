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
	if err := s.set.Init(); err != nil {
		return err
	}
	// menu init
	if err := s.menu.Init(); err != nil {
		return err
	}
	// store init
	if err := s.store.Init(); err != nil {
		return err
	}
	return nil
}

func (gs *SetService) ServiceShutdown() error {
	return nil
}

// setting
func (s *SetService) GetConfig() vo.SettingVo {
	return s.set.GetSetting()
}

func (s *SetService) UpdateConfig(config vo.UpdateSettingVo) error {
	return s.set.UpdateSetting(config)
}

// menu
func (s *SetService) GetMenuList() []vo.MenuVo {
	return s.menu.GetMenuList()
}

func (s *SetService) UpdateMenu(menu vo.UpdateMenuVo) error {
	return s.UpdateMenu(menu)
}

// store
func (s *SetService) GetStore(key string) string {
	return s.store.GetStore(key)
}

func (s *SetService) SetStore(key, value string) error {
	return s.store.SetStore(key, value)
}

func (s *SetService) RemoveStore(key string) error {
	return s.store.DeleteStore(key)
}

func (s *SetService) ClearStore() error {
	return s.store.ClearStore()
}
