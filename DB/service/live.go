package service

import (
	"context"
	_ "embed"
	"errors"
	"pql/DB"
	"pql/DB/model"
	"pql/pkg/vo"

	"gorm.io/gorm"
)

type LiveService struct {
	db  *DB.Sqlite
	ctx context.Context
}

func NewLiveService(db *DB.Sqlite, ctx context.Context) *LiveService {
	return &LiveService{
		db:  db,
		ctx: ctx,
	}
}

func (g *LiveService) GetLiveList(params vo.LiveParams) []vo.LiveVo {
	var liveList = make([]vo.LiveVo, 0)
	query := g.db.WithContext(g.ctx).
		Model(&model.Live{}).
		Select("*").
		Order("is_special ASC, sort ASC, created_at DESC")
	if params.Type != "" {
		query.Where("type = ?", params.Type)
	}
	if params.IsSpecial != "" {
		query.Where("is_special = ?", params.IsSpecial)
	}
	query.Find(&liveList)
	return liveList
}

func (g *LiveService) GetLiveInfo(roomId string, roomType string) *vo.LiveVo {
	var live *vo.LiveVo
	err := g.db.WithContext(g.ctx).Model(&model.Live{}).
		Where("room_id = ?", roomId).
		Where("type = ?", roomType).
		Take(&live).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return live
}

func (g *LiveService) CreateLive(live vo.CreateLiveVo) error {
	if live := g.GetLiveInfo(live.RoomId, live.Type); live != nil {
		return errors.New("房间已存在！")
	}
	result := g.db.WithContext(g.ctx).
		Model(&model.Live{}).
		Create(&model.Live{
			RoomId:    live.RoomId,
			Type:      live.Type,
			IsSpecial: live.IsSpecial,
		})
	return result.Error
}

func (g *LiveService) UpdateLive(live vo.UpdateLiveVo) error {
	result := g.db.WithContext(g.ctx).
		Model(&model.Live{}).
		Select("sort", "is_special").
		Where("id = ?", live.Id).
		Updates(model.Live{
			Sort:      live.Sort,
			IsSpecial: live.IsSpecial,
		})
	return result.Error
}

func (p *LiveService) DeleteLive(ids []int) error {
	return p.db.WithContext(p.ctx).Model(&model.Live{}).Unscoped().Delete(&model.Live{}, ids).Error
}
