package types

type WindowState struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type WindowOptions struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Title string `json:"title"`
}

type PageChange struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type WindowTheme struct {
	Type  string `json:"type"`
	Theme int    `json:"theme"`
}
