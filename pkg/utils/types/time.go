package types

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return
	}
	t.Time, err = time.Parse(time.TimeOnly, string(data))
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte("\"" + t.Format(time.TimeOnly) + "\""), nil
}

// 转换为数据库值
func (t Time) Value() (driver.Value, error) {

	if t.IsZero() {
		return nil, nil
	}

	return t.Time, nil
}

// 数据库值转换为Time
func (t *Time) Scan(value interface{}) error {

	if val, ok := value.(time.Time); ok {
		*t = Time{Time: val}
		return nil
	}

	if val, ok := value.([]byte); ok {
		v := string(val)
		if v == "" {
			return nil
		}
		tv, err := time.Parse(time.TimeOnly, v)
		if err != nil {
			return nil
		}
		*t = Time{Time: tv}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}
