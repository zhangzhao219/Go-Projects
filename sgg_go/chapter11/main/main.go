package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p *Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p *Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

func (c *Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c *Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (c *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (s *Stu) Say() {
	fmt.Println(s.Name)
}

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (h *HeroSlice) Len() int {
	return len(*h)
}

func (h *HeroSlice) Less(i, j int) bool {
	if (*h)[i].Name < (*h)[j].Name {
		return true
	} else {
		return (*h)[i].Age < (*h)[j].Age
	}
}

func (h *HeroSlice) Swap(i, j int) {
	temp := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = temp
}

func main() {
	computer := &Computer{}
	phone := &Phone{}
	camera := &Camera{}
	computer.Working(phone)
	computer.Working(camera)

	var stu = &Stu{"fdsf"}
	var b AInterface = stu
	b.Say()
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("%d", rand.Intn(100)),
			Age:  rand.Intn(10),
		}
		heros = append(heros, hero)
	}
	for _, v := range heros {
		fmt.Println(v)
	}
	fmt.Println()
	sort.Sort(&heros)
	for _, v := range heros {
		fmt.Println(v)
	}
}
