package model

type Parser struct {
	Type        string `json:"type"`
	Source      string `json:"source"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	SubType     string `json:"subType"`
	HomePage    string `json:"homePage"`
	Token       string `json:"token"`
	BaseModel
}

func (*Parser) TableName() string {
	return "user_parser"
}
