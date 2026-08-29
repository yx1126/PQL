package vo

import "pql/DB/model"

type BaseSettingVo struct {
	CloseBehavior        int    `json:"closeBehavior"`
	Theme                int    `json:"theme"`
	Lang                 string `json:"lang"`
	VideoDetailTabActive string `json:"videoDetailTabActive"`
	VideoDetailGrid      string `json:"videoDetailGrid"`
	VideoDetailSort      string `json:"videoDetailSort"`
	VideoSourceType      string `json:"videoSourceType"`
	AnimeSourceType      string `json:"animeSourceType"`
	LiveShowType         string `json:"liveShowType"`
	LiveSpecialShowType  string `json:"liveSpecialShowType"`
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
