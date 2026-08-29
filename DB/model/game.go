package model

type Game struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// 启动指令
	Command string `json:"command"`
	// 默认启动指令
	CustomCommand string `json:"customCommand"`
	// 安装目录
	InstallFolder string `json:"installFolder"`
	// mods 安装目录
	ModsFolder string `json:"modsFolder"`
	// 存档
	SaveFolder string `json:"saveFolder"`
	// 大小
	Size int `json:"size" gorm:"default:0"`
	// mod大小
	ModsSize int `json:"modsSize" gorm:"default:0"`
	// 新弹窗
	IsSupportOpenWindow int `json:"isSupportOpenWindow" gorm:"default:0"`
	// 路径
	Path string `json:"path"`
	// 是否固定
	IsFixed int `json:"isFixed" gorm:"default:0"`
	// 隐藏
	Hidden int `json:"hidden" gorm:"default:0"`
	BaseModel
}

func (*Game) TableName() string {
	return "user_game"
}
