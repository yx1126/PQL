package vo

import (
	"errors"
	"pql/DB/model"
	"strings"
)

type BaseStore struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *BaseStore) Valid() error {
	if strings.TrimSpace(s.Key) == "" {
		return errors.New("Key 不能为空")
	}
	return nil
}

type StoreVo struct {
	BaseStore
	model.BaseModel
}

type CreateStoreVo struct {
	BaseStore
}

type UpdateStoreVo struct {
	BaseStore
}
