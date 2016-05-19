package easyinterface

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestIsArray(t *testing.T) {
	e := FromJSON(`[1, 2, 3]`)
	if !e.IsArray() {
		t.Error("Expecting array")
	}
	if e.IsObject() {
		t.Error("Not expecting object")
	}
}

func TestIsObject(t *testing.T) {
	e := FromJSON(`{"number": [1, 2, 3]}`)
	if e.IsArray() {
		t.Error("Not expecting array")
	}
	if !e.IsObject() {
		t.Error("Expecting object")
	}
}

func TestGetArray(t *testing.T) {
	e := FromJSON(`["Go", "PHP"]`)

	var v *EasyInterface

	v = e.Get("0")
	if v.ToString() != "Go" {
		t.Error(fmt.Sprintf("Expecting Go, got %#v", v))
	}

	if !e.Get("2").IsNil() {
		t.Error(fmt.Sprintf("Expecting true value, got %#v", v))
	}

	if !e.Get("pre-school").IsNil() {
		t.Error(fmt.Sprintf("Expecting true value, got %#v", v))
	}
}

func TestGetObject(t *testing.T) {
	e := FromJSON(`{"name":"yeka","number":13,"skills": ["Go", "PHP"]}`)
	if e.Get("name").ToString() != "yeka" {
		t.Error("Expecting yeka")
	}
	if e.Get("number").ToInt() != 13 {
		t.Error("Expecting 13")
	}
	if e.Get("skills[0]").ToString() != "Go" {
		t.Error("Expecting Go")
	}

	if !e.Get("number[2]").IsNil() {
		t.Error("Expecting nil value")
	}
	if !e.Get("none").IsNil() {
		t.Error(fmt.Sprintf("Expecting true"))
	}
//
//	s := e.Get("none").ToString()
//	if s != "" {
//		t.Error(fmt.Sprintf("Expecting empty string, got %#v", s))
//	}
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
	assert(t, 2, (&EasyInterface{json.Number("2")}).ToInt())
	assert(t, 0, (&EasyInterface{[]byte(`hello`)}).ToInt())
}

func TestToStringArray(t *testing.T) {
	a := FromJSON(`["Go","PHP"]`).ToStringArray()
	assert(t, 2, len(a))
	assert(t, "Go", a[0])

	b := FromJSON(`[1, 2, 3]`).ToStringArray()
	assert(t, 3, len(b))
	assert(t, "1", b[0])
	assert(t, "2", b[1])
}

func TestToIntArray(t *testing.T) {
	a := FromJSON(`["Go","PHP", "12"]`).ToIntArray()
	assert(t, 3, len(a))
	assert(t, 0, a[0])
	assert(t, 0, a[1])
	assert(t, 12, a[2])

	b := FromJSON(`[1, 2]`).ToIntArray()
	assert(t, 2, len(b))
	assert(t, 1, b[0])
	assert(t, 2, b[1])
}

func assert(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Error(fmt.Sprintf("Expecting %#v, got %#v\n", expected, actual))
	}
}