package main

import (
	"encoding/json"
	"fmt"
)

func jsonDemo() {
	// struct -> json（序列化）
	s := Student{Name: "小明", Score: 95}
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("出错：", err)
	}

	fmt.Println(string(data)) // {"Name": "小明", "Score": 95}

	// Json -> struct(反序列化)
	jsonStr := `{"Name":"小红","Score":88}`
	var s2 Student
	err = json.Unmarshal([]byte(jsonStr), &s2)
	if err != nil {
		fmt.Println("出错：", err)
		return
	}
	fmt.Println(s2.Name, s2.Score)
}

//   ┌───────────────┬───────────┬──────────┐
//   │     方向       │   方法     │   类比   │
//   ├───────────────┼───────────┼──────────┤
//   │ struct → JSON │ Marshal   │ 打包快递   │
//   ├───────────────┼───────────┼──────────┤
//   │ JSON → struct │ Unmarshal │ 拆快递    │
//   └───────────────┴───────────┴──────────┘
