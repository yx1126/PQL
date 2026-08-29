package service

import (
	"context"
	"pql/DB"
	"pql/DB/model"
	"pql/DB/scopes"
	"pql/pkg/vo"
)

type MenuService struct {
	db  *DB.Sqlite
	ctx context.Context
}

func NewMenuService(db *DB.Sqlite, ctx context.Context) *MenuService {
	return &MenuService{
		db:  db,
		ctx: ctx,
	}
}

func (m *MenuService) Init() error {
	var menuCount int64
	m.db.WithContext(m.ctx).Model(&model.Menu{}).Count(&menuCount)
	if menuCount <= 0 {
		menu := []model.Menu{
			{Name: "首页", Icon: "home", Path: "/home", Size: 16, Hidden: 0},
			{Name: "游戏", Icon: "games", Path: "/home/games", Size: 16, Hidden: 0},
			{Name: "直播", Icon: "live", Path: "/home/live", Size: 20, Hidden: 0},
			{Name: "音乐", Icon: "music", Path: "/home/music", Size: 16, Hidden: 0},
			{Name: "视频", Icon: "video", Path: "/home/video", Size: 16, Hidden: 0},
			{Name: "动漫", Icon: "anime", Path: "/home/anime", Size: 20, Hidden: 0},
		}
		if err := m.db.WithContext(m.ctx).Model(&model.Menu{}).Create(&menu).Error; err != nil {
			return err
		}
	}
	return nil
}

func (m *MenuService) GetMenuList() []vo.MenuVo {
	var menuList = make([]vo.MenuVo, 0)
	m.db.WithContext(m.ctx).Model(&model.Menu{}).Select("*").Find(&menuList)
	return menuList
}

func (m *MenuService) UpdateMenu(menu vo.UpdateMenuVo) error {
	result := m.db.WithContext(m.ctx).
		Model(&model.Menu{}).
		Scopes(scopes.UpdateOmitScope()).
		Where("id = ?", menu.Id).
		Updates(&model.Menu{
			Name:   menu.Name,
			Icon:   menu.Icon,
			Path:   menu.Path,
			Size:   menu.Size,
			Hidden: menu.Hidden,
		})
	return result.Error
}
