package vo

import "pql/DB/model"

type BaseSettingVo struct {
	CloseBehavior        int    `json:"closeBehavior"`
	Theme                int    `json:"theme"`
	ColorTheme           string `json:"colorTheme"`
	Lang                 string `json:"lang"`
	VideoDetailTabActive string `json:"videoDetailTabActive"`
	VideoDetailGrid      string `json:"videoDetailGrid"`
	VideoDetailSort      string `json:"videoDetailSort"`
	VideoSourceType      string `json:"videoSourceType"`
	AnimeSourceType      string `json:"animeSourceType"`
	LiveShowType         string `json:"liveShowType"`
	HeaderShowTheme      string `json:"headerShowTheme"`
	AnimeWeeklyType      string `json:"animeWeeklyType"`
}

type SettingVo struct {
	BaseSettingVo
	model.BaseModel
}

type UpdateSettingVo struct {
	BaseSettingVo
	BaseVo
}
