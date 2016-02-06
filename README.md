# goeasy

##1. EasyInterface

Help you convert `interface{}` data easily

Usage example:
```go
package main

import (
	"fmt"
	"encoding/json"
	"github.com/yeka/goeasy/easyinterface"
)

func main() {
	var i interface{}
	json.Unmarshal([]byte(`{"name":"yeka","skills":[{"lang": "Go"}, {"lang":"PHP"}]}`), &i)
	e := easyinterface.EasyInterface{i}

	// Now access the value easily using PHP-associative-array-like syntax
	fmt.Println(e.Get("name").ToString())
	fmt.Println(e.Get("skills[0][lang]").ToString())
}
```
