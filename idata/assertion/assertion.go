package assertion

import (
	"bytes"
	"fmt"
	"reflect"

	. "github.com/mataharimall/api-seller/tests/functional/idata"
)

func ShouldBeJSONAndHave(actual interface{}, expected ...interface{}) string {
	if len(expected) != 2 {
		return "Must have key/value arguments"
	}

	var b []byte
	if v, ok := actual.([]byte); ok {
		b = v
	} else if v, ok := actual.(*bytes.Buffer); ok {
		b = v.Bytes()
	} else {
		return "Actual types should be []byte or *bytes.Buffer"
	}
	json := FromJSON(b)

	if json.GetString(expected[0].(string)) != expected[1].(string) {
		return fmt.Sprintf(`Expected: "%v" = %v`+"\nActual: %v", expected[0], expected[1], string(b))
	}

	return ""
}

func ShouldBeJSONAndCount(actual interface{}, expected ...interface{}) string {
	if len(expected) != 2 {
		return "Must have key/number of result"
	}

	var b []byte
	if v, ok := actual.([]byte); ok {
		b = v
	} else if v, ok := actual.(*bytes.Buffer); ok {
		b = v.Bytes()
	} else {
		return "Actual types should be []byte or *bytes.Buffer"
	}
	json := FromJSON(b)

	if json.Count(expected[0].(string)) != expected[1].(int) {
		s := ""
		if expected[1].(int) > 1 {
			s = "s"
		}
		return fmt.Sprintf(`Expected: "%v" have %v item%v `+"\nActual: %v", expected[0], expected[1], s, string(b))
	}

	return ""
}

func ShouldBeJSONAndContain(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 || reflect.TypeOf(expected[0]).String() != "string" {
		return "Must have JSON string as expectation"
	}

	var b []byte
	if v, ok := actual.([]byte); ok {
		b = v
	} else if v, ok := actual.(*bytes.Buffer); ok {
		b = v.Bytes()
	} else {
		return "Actual types should be []byte or *bytes.Buffer"
	}
	json := FromJSON(b)
	je := FromJSON([]byte(expected[0].(string)))

	comparison := compare(json, je, "")
	if comparison != "" {
		return fmt.Sprintf(`Expected: %v`+"\nActual: %v", comparison, string(b))
	}

	return ""
}

func ShouldBeJSONAndNotContain(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 || reflect.TypeOf(expected[0]).String() != "string" {
		return "Must have JSON string as expectation"
	}

	var b []byte
	if v, ok := actual.([]byte); ok {
		b = v
	} else if v, ok := actual.(*bytes.Buffer); ok {
		b = v.Bytes()
	} else {
		return "Actual types should be []byte or *bytes.Buffer"
	}
	json := FromJSON(b)
	je := FromJSON([]byte(expected[0].(string)))

	comparison := compare(json, je, "")
	if comparison == "" {
		return fmt.Sprintf(`Expected: %v`+"\nActual: %v", comparison, string(b))
	}

	return ""
}

func compare(actual, expected *IData, key string) string {
	if expected.IsObject(key) && actual.IsObject(key) {
		m, _ := expected.Get(key)
		for i := range m.(map[string]interface{}) {
			newkey := i
			if key != "" {
				newkey = key + "." + newkey
			}
			if res := compare(actual, expected, newkey); res != "" {
				return res
			}
		}
		return ""
	} else if expected.IsArray(key) && actual.IsArray(key) {
		me, _ := expected.Get(key)
		ma, _ := actual.Get(key)
		for _, ei := range me.([]interface{}) {
			res := ""
			found := false
			for _, ai := range ma.([]interface{}) {
				if res = compare(New(ai), New(ei), ""); res == "" {
					found = true
				}
			}
			if !found {
				return fmt.Sprintf(`%v[].%v`, key, res)
			}
		}
		return ""
	}
	e, _ := expected.Get(key)
	a, _ := actual.Get(key)
	if e != a {
		return fmt.Sprintf(`%v = %v`, key, e)
	}
	return ""
}
