package service

import (
	"context"
	"pql/DB"
	"pql/DB/model"
	"pql/pkg/vo"

	"gorm.io/gorm/clause"
)

type AuthService struct {
	db  *DB.Sqlite
	ctx context.Context
}

func NewAuthService(db *DB.Sqlite, ctx context.Context) *AuthService {
	return &AuthService{
		db:  db,
		ctx: ctx,
	}
}

func (s *AuthService) GetAuthListist() []vo.AuthVo {
	var authList = make([]vo.AuthVo, 0)
	s.db.WithContext(s.ctx).Model(&model.Auth{}).Select("*").Find(&authList)
	return authList
}

func (s *AuthService) GetAuth(typee string) (*vo.AuthVo, error) {
	var auth vo.AuthVo
	if err := s.db.WithContext(s.ctx).
		Model(&model.Auth{}).
		Where("type = ?", typee).
		Take(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func (s *AuthService) SaveAuth(auth vo.SaveAuthVo) error {
	result := s.db.WithContext(s.ctx).
		Model(&model.Auth{}).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "type"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"token",
				"expires_in",
				"expires_time",
				"refresh_token",
				"scope",
			}),
		}).
		Create(&model.Auth{
			Type:         auth.Type,
			Token:        auth.Token,
			ExpiresIn:    auth.ExpiresIn,
			ExpiresTime:  auth.ExpiresTime,
			RefreshToken: auth.RefreshToken,
			Scope:        auth.Scope,
		})
	return result.Error
}

func (s *AuthService) DeleteStore(typee string) error {
	return s.db.WithContext(s.ctx).
		Model(&model.Auth{}).
		Unscoped().
		Where("type = ?", typee).
		Delete(&model.Auth{}).Error
}
