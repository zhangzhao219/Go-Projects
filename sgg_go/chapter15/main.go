package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"aaaaaaaaname"`
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() []byte {
	monster := &Monster{
		Name:     "fdsfds",
		Age:      45,
		Birthday: "fdsdsfds",
		Sal:      0.62,
		Skill:    "fdsfdsds",
	}
	content, _ := json.Marshal(monster)
	var m1 Monster
	json.Unmarshal(content, &m1)
	fmt.Println(m1)
	return content
}

func testMap() []byte {
	a := make(map[string]int)
	a["fdsfds"] = 0
	a["erewr"] = 1
	content, _ := json.Marshal(a)
	json.Unmarshal(content, &a)
	fmt.Println(a)
	return content
}

func testSlice() []byte {
	var slice []map[string]int
	a := make(map[string]int)
	a["fdsfds"] = 0
	a["erewr"] = 1
	slice = append(slice, a)
	b := make(map[string]int)
	b["fddsfdsdsfsfds"] = 0
	b["eredsfdsfdsfdswr"] = 1
	slice = append(slice, b)
	content, _ := json.Marshal(slice)
	json.Unmarshal(content, &slice)
	fmt.Println(slice)
	return content
}

func main() {
	fmt.Println(string(testStruct()))
	fmt.Println(string(testMap()))
	fmt.Println(string(testSlice()))
}
