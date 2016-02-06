package easyinterface

import (
	"strconv"
	"regexp"
)

type EasyInterface struct {
	Interface interface{}
}

func (this *EasyInterface) ToInt() int {
	switch v := this.Interface.(type) {
	case string:
		res, _ := strconv.Atoi(v)
		return res
	case int:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	}
	return 0
}

func (this *EasyInterface) ToString() string {
	switch v := this.Interface.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case float32:
		return strconv.FormatInt(int64(v), 10)
	case float64:
		return strconv.FormatInt(int64(v), 10)
	}
	return ""
}

func (this *EasyInterface) ToObject() map[string]EasyInterface {
	data := map[string]EasyInterface{}
	switch m := this.Interface.(type) {
	case map[string]interface{}:
		for k, v := range m {
			data[k] = EasyInterface{v}
		}
		break
	}
	return data
}

func (this *EasyInterface) ToArray() []EasyInterface {
	data := make([]EasyInterface, 0)
	switch v := this.Interface.(type) {
	case []interface{}:
		for _, i := range v {
			data = append(data, EasyInterface{i})
		}
		break
	}
	return data
}

func (this *EasyInterface) ToIntArray() []int {
	data := make([]int, 0)
	switch v := this.Interface.(type) {
	case []interface{}:
		for _, i := range v {
			ii := EasyInterface{i}
			data = append(data, ii.ToInt())
		}
		break
	}
	return data
}

func (this *EasyInterface) ToStringArray() []string {
	data := make([]string, 0)
	switch v := this.Interface.(type) {
	case []interface{}:
		for _, i := range v {
			ii := EasyInterface{i}
			data = append(data, ii.ToString())
		}
		break
	}
	return data
}

func (this *EasyInterface) IsArray() bool {
	switch this.Interface.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}

func (this *EasyInterface) IsObject() bool {
	switch this.Interface.(type) {
	case map[string]interface{}:
		return true
	default:
		return false
	}
}

func (this *EasyInterface) Get(pattern string) *EasyInterface {
	r1, _ := regexp.Compile(`^([a-z0-9]+)(\[([a-z0-9]+)\])*$`)
	if !r1.MatchString(pattern) {
		return nil
	}

	var result *EasyInterface

	f1 := r1.FindStringSubmatch(pattern)[1]
	result = this.getinside(f1)

	r2, _ := regexp.Compile(`\[([a-z0-9]+)\]`)
	f2 := r2.FindAllStringSubmatch(pattern, -1)

	for _, v := range f2 {
		result = result.getinside(v[1])
	}

	return result
}

func (this *EasyInterface) getinside(key string) *EasyInterface {
	var result EasyInterface
	var ok bool

	if this.IsObject() {
		result, ok = this.ToObject()[key]
		if !ok {
			return nil
		}
	} else if this.IsArray() {
		index, _ := strconv.Atoi(key)
		array := this.ToArray()
		if index >= len(array) {
			return nil
		}
		result = array[index]
	} else {
		return nil
	}

	return &result
}