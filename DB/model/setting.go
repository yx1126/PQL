package model

// 基础设置
type Setting struct {

	// 关闭按钮行为 隐藏窗口:0 直接推出:1
	CloseBehavior int `json:"closeBehavior" gorm:"default:0"`
	// 主题 暗黑:0 亮色:1 跟随系统:2
	Theme int `json:"theme" gorm:"default:0"`
	// 语言
	Lang string `json:"lang"`

	// 视频详情页
	// Tab切换
	VideoDetailTabActive string `json:"videoDetailTabActive" gorm:"default:info"`
	// 列表布局
	VideoDetailGrid string `json:"videoDetailGrid" gorm:"default:default"`
	// 排序规则
	VideoDetailSort string `json:"videoDetailSort" gorm:"default:asc"`
	// 源
	VideoSourceType string `json:"videoSourceType"`
	AnimeSourceType string `json:"animeSourceType"`

	// 直播
	// 显示类型
	LiveShowType        string `json:"liveShowType" gorm:"default:all"`
	LiveSpecialShowType string `json:"liveSpecialShowType" gorm:"default:all"`

	// 动漫
	// 周番显示类型
	AnimeWeeklyType string `json:"animeWeeklyType"`
	BaseModel
}

func (*Setting) TableName() string {
	return "user_set"
}
