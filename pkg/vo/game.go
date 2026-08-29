package vo

import "pql/DB/model"

type GameVo struct {
	Name                string `json:"name"`
	Type                string `json:"type"`
	Command             string `json:"command"`
	CustomCommand       string `json:"customCommand"`
	InstallFolder       string `json:"installFolder"`
	ModsFolder          string `json:"modsFolder"`
	SaveFolder          string `json:"saveFolder"`
	Size                int    `json:"size"`
	ModsSize            int    `json:"modsSize"`
	IsSupportOpenWindow int    `json:"isSupportOpenWindow"`
	Path                string `json:"path"`
	IsFixed             int    `json:"isFixed"`
	Hidden              int    `json:"hidden"`
	model.BaseModel
}

type UpdateGameVo struct {
	CustomCommand string `json:"customCommand"`
	InstallFolder string `json:"installFolder"`
	ModsFolder    string `json:"modsFolder"`
	SaveFolder    string `json:"saveFolder"`
	Path          string `json:"path"`
	IsFixed       int    `json:"isFixed"`
	Hidden        int    `json:"hidden"`
	BaseVo
}
