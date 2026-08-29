package service

import (
	"context"
	"pql/DB"
	"pql/DB/model"
	"pql/pkg/constant"
	"pql/pkg/parse/live"

	"gorm.io/gorm/clause"
)

type StoreService struct {
	db  *DB.Sqlite
	ctx context.Context
}

func NewStoreService(db *DB.Sqlite, ctx context.Context) *StoreService {
	return &StoreService{
		db:  db,
		ctx: ctx,
	}
}

func (s *StoreService) Init() error {
	if err := s.initDouyin(); err != nil {
		return err
	}
	return nil
}

// 抖音接口鉴权 ttwid
func (s *StoreService) initDouyin() error {
	var ttwidCount int64
	s.db.WithContext(s.ctx).Model(&model.Store{}).Where("key = ?", constant.Ttwid).Count(&ttwidCount)
	if ttwidCount <= 0 {
		if err := s.db.WithContext(s.ctx).
			Model(&model.Store{}).
			Create(&model.Store{
				Key:   constant.Ttwid,
				Value: live.Ttwid,
			}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *StoreService) GetStore(key string) string {
	var value string
	if err := s.db.WithContext(s.ctx).
		Model(&model.Store{}).
		Select("value").
		Where("key = ?", key).
		Pluck("value", &value).Error; err != nil {
		return ""
	}
	return value
}

func (s *StoreService) SetStore(key, value string) error {
	result := s.db.WithContext(s.ctx).
		Model(&model.Store{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"value"}),
		}).
		Create(&model.Store{
			Key:   key,
			Value: value,
		})
	return result.Error
}

func (s *StoreService) DeleteStore(key string) error {
	return s.db.WithContext(s.ctx).Model(&model.Store{}).Unscoped().Where("key = ?", key).Delete(&model.Store{}).Error
}

func (s *StoreService) ClearStore() error {
	return s.db.WithContext(s.ctx).Model(&model.Store{}).Unscoped().Delete(&model.Store{}).Error
}
