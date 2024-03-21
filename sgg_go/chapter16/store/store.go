package store

import (
	"bufio"
	"encoding/json"
	"os"
)

type monster struct {
	Name  string
	Age   int
	Skill string
}

func Newmonster(name string, age int, skill string) *monster {
	return &monster{
		Name:  name,
		Age:   age,
		Skill: skill,
	}
}

func (m *monster) Store(fileName string) bool {
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	writer := bufio.NewWriter(file)
	str, _ := json.Marshal(m)
	writer.WriteString(string(str))
	writer.Flush()
	return true
}

func (m *monster) ReStore(fileName string) *monster {
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0666)
	reader := bufio.NewReader(file)
	str, _ := reader.ReadString('\n')
	var data []byte = []byte(str)
	json.Unmarshal(data, &m)
	return m
}
