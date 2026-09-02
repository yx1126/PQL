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

func (a *AuthService) Init() error {
	var count int64
	a.db.WithContext(a.ctx).Model(&model.Auth{}).Count(&count)
	if count <= 0 {
		auths := []model.Auth{
			{Name: "百度网盘", Type: "baidu", Icon: "baidu"},
		}
		if err := a.db.WithContext(a.ctx).Model(&model.Auth{}).Create(&auths).Error; err != nil {
			return err
		}
	}
	return nil
}

func (a *AuthService) GetAuthListist() []vo.AuthVo {
	var authList = make([]vo.AuthVo, 0)
	a.db.WithContext(a.ctx).Model(&model.Auth{}).Select("*").Find(&authList)
	return authList
}

func (a *AuthService) GetAuth(typee string) (*vo.AuthVo, error) {
	var auth vo.AuthVo
	if err := a.db.WithContext(a.ctx).
		Model(&model.Auth{}).
		Where("type = ?", typee).
		Take(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func (a *AuthService) SaveAuth(auth vo.SaveAuthVo) error {
	result := a.db.WithContext(a.ctx).
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

func (a *AuthService) DeleteAuth(typee string) error {
	return a.SaveAuth(vo.SaveAuthVo{
		BaseAuth: vo.BaseAuth{
			Type:         typee,
			Token:        "",
			ExpiresIn:    0,
			ExpiresTime:  "",
			RefreshToken: "",
			Scope:        "",
		},
	})
}
