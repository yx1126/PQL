package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"pql/pkg/parse/video"
	"pql/pkg/vo"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type ParserService struct {
	*ServiceContext
}

func NewParserService(sc *ServiceContext) *ParserService {
	return &ParserService{
		ServiceContext: sc,
	}
}

// Initialisation
func (s *ParserService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {

	return nil
}

func (gs *ParserService) ServiceShutdown() error {
	return nil
}

func (p *ParserService) GetParserList() []vo.ParserVo {
	return p.parser.GetParserList()
}

func (p *ParserService) CreateParsers(source string) (vo.Result[video.ParseOption], error) {
	result := vo.Result[video.ParseOption]{}
	parserList, err := video.Parse(source)
	if err != nil {
		return result, err
	}
	var successCount, failureCount int

	successList := []video.ParseOption{}

	for _, v := range parserList {
		if err := v.Validator(); err != nil {
			failureCount++
			continue
		}
		rawSource, err := json.Marshal(v)
		if err != nil {
			failureCount++
			continue
		}
		if err := p.parser.CreateParser(vo.CreateParserVo{
			BaseParserVo: vo.BaseParserVo{
				Type:        v.Type,
				Source:      string(rawSource),
				Version:     v.Version,
				Author:      v.Author,
				Description: v.Description,
				Icon:        v.Icon,
				SubType:     v.SubType,
				HomePage:    v.HomePage,
			},
		}); err != nil {
			failureCount++
			continue
		}
		successCount++
		successList = append(successList, v)
	}
	result.Message = fmt.Sprintf("导入成功%d条，失败%d条。", successCount, failureCount)
	result.Data = successList
	return result, nil
}

func (p *ParserService) UpdateToken(id int, token string) error {
	return p.parser.UpdateToken(id, token)
}

func (p *ParserService) DeleteParser(ids []int) error {
	if len(ids) <= 0 {
		return errors.New("请选择要删除的数据")
	}
	return p.parser.DeleteParser(ids)
}
