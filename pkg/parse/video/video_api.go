package video

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func Parse(source string) ([]ParseOption, error) {
	parserList := make([]ParseOption, 0)
	if strings.TrimSpace(source) == "" {
		return parserList, errors.New("格式不正确")
	}
	if err := json.Unmarshal([]byte(source), &parserList); err != nil {
		var data ParseOption
		if err := json.Unmarshal([]byte(source), &data); err == nil {
			parserList = append(parserList, data)
		} else {
			fmt.Print("对象解析", err)
		}
	} else {
		fmt.Print("列表解析", err)
	}
	if len(parserList) <= 0 {
		return parserList, errors.New("格式不正确")
	}
	return parserList, nil
}
