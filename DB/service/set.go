package service

import (
	"context"
	"pql/DB"
	"pql/DB/model"
	"pql/DB/scopes"
	"pql/pkg/vo"
)

type SettingService struct {
	db  *DB.Sqlite
	ctx context.Context
}

func NewSettingService(db *DB.Sqlite, ctx context.Context) *SettingService {
	return &SettingService{
		db:  db,
		ctx: ctx,
	}
}

func (s *SettingService) Init() error {
	var setCount int64
	s.db.WithContext(s.ctx).Model(&model.Setting{}).Count(&setCount)
	if setCount <= 0 {
		sets := model.Setting{
			CloseBehavior:        0,
			Theme:                0,
			Lang:                 "zh-cn",
			VideoDetailTabActive: "info",
			VideoDetailGrid:      "default",
			VideoDetailSort:      "asc",
			VideoSourceType:      "",
			AnimeSourceType:      "",
			LiveShowType:         "default",
			LiveSpecialShowType:  "all",
			AnimeWeeklyType:      "cn",
		}
		if err := s.db.WithContext(s.ctx).Model(&model.Setting{}).Create(&sets).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *SettingService) GetSetting() vo.SettingVo {
	var config vo.SettingVo
	s.db.WithContext(s.ctx).Model(&model.Setting{}).Select("*").First(&config)
	return config
}

func (s *SettingService) UpdateSetting(config vo.UpdateSettingVo) error {
	result := s.db.WithContext(s.ctx).
		Model(&model.Setting{}).
		Scopes(scopes.UpdateOmitScope()).
		Where("id = ?", config.Id).
		Updates(&model.Setting{
			CloseBehavior:        config.CloseBehavior,
			Theme:                config.Theme,
			Lang:                 config.Lang,
			VideoDetailTabActive: config.VideoDetailTabActive,
			VideoDetailGrid:      config.VideoDetailGrid,
			VideoDetailSort:      config.VideoDetailSort,
			VideoSourceType:      config.VideoSourceType,
			AnimeSourceType:      config.AnimeSourceType,
			LiveShowType:         config.LiveShowType,
			LiveSpecialShowType:  config.LiveSpecialShowType,
			AnimeWeeklyType:      config.AnimeWeeklyType,
		})
	return result.Error
}
