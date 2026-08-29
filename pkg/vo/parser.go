package vo

import (
	"pql/DB/model"
)

type BaseParserVo struct {
	Type        string `json:"type"`
	Source      string `json:"source"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	SubType     string `json:"subType"`
	HomePage    string `json:"homePage"`
	Token       string `json:"token"`
}

type ParserVo struct {
	BaseParserVo
	model.BaseModel
}

type CreateParserVo struct {
	BaseParserVo
}

type UpdateParserVo struct {
	BaseParserVo
	BaseVo
}
