package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	value, ok := users[name]
	if ok {
		value["password"] = "888888"
	} else {
		value2 := make(map[string]string)
		value2["nickname"] = name
		value2["password"] = "888889"
		users[name] = value2
	}
}

func main() {
	var a = make(map[string]string)
	fmt.Println(a)
	student := map[int]map[string]string{
		1: {
			"name": "aaa",
			"sex":  "男",
		},
		2: {
			"name": "bbb",
			"sex":  "男",
		},
		3: {
			"name": "ccc",
			"sex":  "女",
		},
	}
	for _, v := range student {
		for k1, v1 := range v {
			fmt.Println(k1, v1)
		}
	}
	delete(student, 3)
	fmt.Println(student)
	val, ok := student[8]
	if ok {
		fmt.Println(val)
	}
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "dsfd"
		monsters[0]["age"] = "edf"
	}
	fmt.Println(monsters)

	users := make(map[string]map[string]string)
	modifyUser(users, "zz")
	fmt.Println(users)
	modifyUser(users, "zz")
	fmt.Println(users)
	modifyUser(users, "zz")
	fmt.Println(users)
	modifyUser(users, "yy")
	fmt.Println(users)
}
