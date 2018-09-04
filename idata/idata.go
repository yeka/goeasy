package idata

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type IData struct {
	data interface{}
}

func New(i interface{}) *IData {
	return &IData{i}
}

func FromJSON(buf []byte) *IData {
	var i interface{}
	json.Unmarshal(buf, &i)
	return &IData{i}
}

func (i *IData) KeyExists(key string) bool {
	_, err := i.Get(key)
	return err == nil
}

func (i *IData) Get(key string) (interface{}, error) {
	if i.data == nil {
		return nil, errors.New("No data")
	}

	var ii = i.data

	if key != "" {
		keys := strings.Split(key, ".")
		for _, k := range keys {
			switch v := ii.(type) {
			case []interface{}:
				ki, err := strconv.Atoi(k)
				if err != nil {
					return nil, err
				}
				if ki >= len(v) {
					return nil, errors.New("Out of range")
				}
				ii = v[ki]
			case map[string]interface{}:
				var ok bool
				ii, ok = v[k]
				if !ok {
					return nil, errors.New("Key not found")
				}
			}
		}
	}

	return ii, nil
}

func (i *IData) GetString(key string) string {
	ii, err := i.Get(key)
	if err == nil {
		switch v := ii.(type) {
		case string:
			return v
		case int:
			return strconv.FormatInt(int64(v), 10)
		case float32:
			return strconv.FormatInt(int64(v), 10)
		case float64:
			return strconv.FormatInt(int64(v), 10)
		}
	}
	return ""
}

func (i *IData) IsArray(key string) bool {
	ii, err := i.Get(key)
	if err == nil {
		switch ii.(type) {
		case []interface{}:
			return true
		}
	}
	return false
}

func (i *IData) IsObject(key string) bool {
	ii, err := i.Get(key)
	if err == nil {
		switch ii.(type) {
		case map[string]interface{}:
			return true
		}
	}
	return false
}

func (i *IData) Count(key string) int {
	ii, err := i.Get(key)
	if err == nil {
		switch v := ii.(type) {
		case map[string]interface{}:
			return len(v)
		case []interface{}:
			return len(v)
		}
	}
	return 0
}
