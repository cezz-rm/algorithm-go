package main

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
)

type People struct {
	Name string `json:"name"`
}

func main() {
	js := `{"name":"11"}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people:", p)
}

var value int32

func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {
			break
		}
	}
}
