package service

import (
	"context"
	"errors"
	"pql/DB"
	"pql/DB/model"
	"pql/pkg/vo"

	"gorm.io/gorm"
)

type ParserService struct {
	db  *DB.Sqlite
	ctx context.Context
}

func NewParserService(db *DB.Sqlite, ctx context.Context) *ParserService {
	return &ParserService{
		db:  db,
		ctx: ctx,
	}
}

func (p *ParserService) GetParserList() []vo.ParserVo {
	var parseList = make([]vo.ParserVo, 0)
	p.db.WithContext(p.ctx).Model(&model.Parser{}).Select("*").Find(&parseList)
	return parseList
}

func (p *ParserService) GetParser(parserType, subType string) (*vo.ParserVo, error) {
	var parser vo.ParserVo
	result := p.db.WithContext(p.ctx).
		Model(&model.Parser{}).
		Where("type = ?", parserType).
		Where("sub_type = ?", subType).
		Take(&parser)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &parser, nil
}

func (p *ParserService) CreateParser(parser vo.CreateParserVo) error {
	tx := p.db.WithContext(p.ctx).Begin()

	var oldParser *vo.ParserVo

	err := tx.Model(&model.Parser{}).
		Where("type = ?", parser.Type).
		Where("sub_type = ?", parser.SubType).
		Take(&oldParser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := tx.Model(&model.Parser{}).
			Create(&model.Parser{
				Type:        parser.Type,
				Source:      parser.Source,
				Version:     parser.Version,
				Author:      parser.Author,
				Description: parser.Description,
				Icon:        parser.Icon,
				SubType:     parser.SubType,
				HomePage:    parser.HomePage,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := tx.Model(&model.Parser{}).
			Select("type", "source", "version", "author", "description", "icon", "sub_type", "home_page").
			Where("id = ?", oldParser.Id).
			Updates(&model.Parser{
				Type:        parser.Type,
				Source:      parser.Source,
				Version:     parser.Version,
				Author:      parser.Author,
				Description: parser.Description,
				Icon:        parser.Icon,
				SubType:     parser.SubType,
				HomePage:    parser.HomePage,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (p *ParserService) UpdateParser(parser vo.UpdateParserVo) error {
	return p.db.WithContext(p.ctx).
		Model(&model.Parser{}).
		Where("id = ?", parser.Id).
		Updates(&model.Parser{
			Type:        parser.Type,
			Source:      parser.Source,
			Version:     parser.Version,
			Author:      parser.Author,
			Description: parser.Description,
			Icon:        parser.Icon,
			SubType:     parser.SubType,
			HomePage:    parser.HomePage,
		}).Error
}

func (p *ParserService) UpdateToken(id int, token string) error {
	return p.db.WithContext(p.ctx).
		Model(&model.Parser{}).
		Where("id = ?", id).
		Updates(&model.Parser{
			Token: token,
		}).Error
}

func (p *ParserService) DeleteParser(ids []int) error {
	return p.db.WithContext(p.ctx).Model(&model.Parser{}).Unscoped().Delete(&model.Parser{}, ids).Error
}

func (p *ParserService) DeleteAll() error {
	return p.db.WithContext(p.ctx).Model(&model.Parser{}).Unscoped().Delete(&model.Parser{}).Error
}
