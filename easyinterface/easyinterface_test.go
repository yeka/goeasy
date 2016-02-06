package easyinterface

import (
	"testing"
//	"github.com/yeka/goeasy/easyinterface"
	"encoding/json"
	"fmt"
)

func createEasyInterfaceFromJSON(s string) *EasyInterface {
	var i interface{}
	json.Unmarshal([]byte(s), &i)
	return &EasyInterface{i}
}

func TestIsArray(t *testing.T) {
	e := createEasyInterfaceFromJSON(`[1, 2, 3]`)
	if !e.IsArray() {
		t.Error("Expecting array")
	}
	if e.IsObject() {
		t.Error("Not expecting object")
	}
}

func TestIsObject(t *testing.T) {
	e := createEasyInterfaceFromJSON(`{"number": [1, 2, 3]}`)
	if e.IsArray() {
		t.Error("Not expecting array")
	}
	if !e.IsObject() {
		t.Error("Expecting object")
	}
}

func TestGetArray(t *testing.T) {
	e := createEasyInterfaceFromJSON(`["Go", "PHP"]`)

	var v *EasyInterface

	v = e.Get("0")
	if v.ToString() != "Go" {
		t.Error(fmt.Sprintf("Expecting nil value, got %#v", v))
	}

	v = e.Get("2")
	if v != nil {
		t.Error(fmt.Sprintf("Expecting nil value, got %#v", v))
	}

	v = e.Get("pre-school")
	if v != nil {
		t.Error(fmt.Sprintf("Expecting nil value, got %#v", v))
	}
}

func TestGetObject(t *testing.T) {
	e := createEasyInterfaceFromJSON(`{"name":"yeka","number":13,"skills": ["Go", "PHP"]}`)
	if e.Get("name").ToString() != "yeka" {
		t.Error("Expecting yeka")
	}
	if e.Get("number").ToInt() != 13 {
		t.Error("Expecting 13")
	}
	if e.Get("skills[0]").ToString() != "Go" {
		t.Error("Expecting Go")
	}

	if e.Get("number[2]") != nil {
		t.Error("Expecting nil value")
	}
	v := e.Get("none")
	if v != nil {
		t.Error(fmt.Sprintf("Expecting nil value, got %#v", v))
	}
}

func TestToString(t *testing.T) {
	assert(t, "2", (&EasyInterface{int(2)}).ToString())
	assert(t, "2", (&EasyInterface{float32(2)}).ToString())
	assert(t, "2", (&EasyInterface{float64(2)}).ToString())
	assert(t, "3", (&EasyInterface{"3"}).ToString())
	assert(t, "", (&EasyInterface{[]byte(`hello`)}).ToString())
}

func TestToInt(t *testing.T) {
	assert(t, 2, (&EasyInterface{int(2)}).ToInt())
	assert(t, 2, (&EasyInterface{float32(2)}).ToInt())
	assert(t, 2, (&EasyInterface{float64(2)}).ToInt())
	assert(t, 2, (&EasyInterface{"2"}).ToInt())
	assert(t, 0, (&EasyInterface{[]byte(`hello`)}).ToInt())
}

func TestToStringArray(t *testing.T) {
	a := createEasyInterfaceFromJSON(`["Go","PHP"]`).ToStringArray()
	assert(t, 2, len(a))
	assert(t, "Go", a[0])

	b := createEasyInterfaceFromJSON(`[1, 2, 3]`).ToStringArray()
	assert(t, 3, len(b))
	assert(t, "1", b[0])
	assert(t, "2", b[1])
}

func TestToIntArray(t *testing.T) {
	a := createEasyInterfaceFromJSON(`["Go","PHP", "12"]`).ToIntArray()
	assert(t, 3, len(a))
	assert(t, 0, a[0])
	assert(t, 0, a[1])
	assert(t, 12, a[2])

	b := createEasyInterfaceFromJSON(`[1, 2]`).ToIntArray()
	assert(t, 2, len(b))
	assert(t, 1, b[0])
	assert(t, 2, b[1])
}

func assert(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Error(fmt.Sprintf("Expecting %#v, got %#v\n", expected, actual))
	}
}