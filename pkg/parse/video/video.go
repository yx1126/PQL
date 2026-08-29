package video

import (
	"errors"
	"strings"
)

type Response struct {
	Codeok      *int   `json:"codeok,omitempty"`
	CodeExpired *int   `json:"codeExpired,omitempty"`
	CodePath    string `json:"codePath,omitempty"`
	DataPath    string `json:"dataPath,omitempty"`
	MessagePath string `json:"messagePath,omitempty"`
}

type LabelValue struct {
	Type      string `json:"type,omitempty"`
	Extra     *[]any `json:"extra,omitempty"`
	Label     string `json:"label,omitempty"`
	Path      string `json:"path,omitempty"`
	Copy      bool   `json:"copy,omitempty"`
	Template  string `json:"template,omitempty"`
	Separator string `json:"separator,omitempty"`
}

type RuleItem struct {
	Url     string `json:"url,omitempty"`
	Method  string `json:"method,omitempty"`
	Options *[]any `json:"options,omitempty"`

	IsAuth bool `json:"isAuth,omitempty"`

	DataPath    string `json:"dataPath,omitempty"`
	ParseType   string `json:"parseType,omitempty"`
	PrimaryPath string `json:"primaryPath,omitempty"`
	NamePath    string `json:"namePath,omitempty"`

	ChildrenPath     string `json:"childrenPath,omitempty"`
	ChildParseType   string `json:"childParseType,omitempty"`
	ChildPrimaryPath string `json:"childPrimaryPath,omitempty"`
	ChildNamePath    string `json:"childNamePath,omitempty"`

	CoverTopPath   string         `json:"coverTopPath,omitempty"`
	CoverLeftPath  string         `json:"coverLeftPath,omitempty"`
	CoverRightPath string         `json:"coverRightPath,omitempty"`
	SrcPath        string         `json:"srcPath,omitempty"`
	Next           map[string]any `json:"next,omitempty"`
	Data           map[string]any `json:"data,omitempty"`

	Paging       bool   `json:"paging,omitempty"`
	TotalPath    string `json:"totalPath,omitempty"`
	RequestCount *int   `json:"requestCount,omitempty"`

	TitlePath       string        `json:"titlePath,omitempty"`
	DetailList      *[]LabelValue `json:"detailList,omitempty"`
	DescriptionList *[]LabelValue `json:"descriptionList,omitempty"`

	Omitempty bool `json:"omitempty,omitempty"`
}

type ParseOption struct {
	Type           string            `json:"type,omitempty"`
	Author         string            `json:"author,omitempty"`
	Version        string            `json:"version,omitempty"`
	Email          string            `json:"email,omitempty"`
	Description    string            `json:"description,omitempty"`
	DefaultSize    *int              `json:"defaultSize,omitempty"`
	DefaultMaxSize *int              `json:"defaultMaxSize,omitempty"`
	Icon           string            `json:"icon,omitempty"`
	HomePage       string            `json:"homePage,omitempty"`
	SubType        string            `json:"subType,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
	Domain         string            `json:"domain,omitempty"`
	Imgdomain      string            `json:"imgdomain,omitempty"`
	Authorization  string            `json:"authorization,omitempty"`

	Response *Response `json:"response,omitempty"`

	LoginApi      *RuleItem `json:"loginApi,omitempty"`
	TypeApi       *RuleItem `json:"typeApi,omitempty"`
	DataApi       *RuleItem `json:"dataApi,omitempty"`
	SearchTypeApi *RuleItem `json:"searchTypeApi,omitempty"`
	SearchApi     *RuleItem `json:"searchApi,omitempty"`
	DetailApi     *RuleItem `json:"detailApi,omitempty"`
	SourceApi     *RuleItem `json:"sourceApi,omitempty"`
	EpisodeApi    *RuleItem `json:"episodeApi,omitempty"`
	PlayUrlApi    *RuleItem `json:"playUrlApi,omitempty"`
	ScheduleApi   *RuleItem `json:"scheduleApi,omitempty"`
}

func (po *ParseOption) Validator() error {
	if strings.TrimSpace(po.Type) == "" ||
		strings.TrimSpace(po.SubType) == "" {
		return errors.New("格式不正确！")
	}
	return nil
}
