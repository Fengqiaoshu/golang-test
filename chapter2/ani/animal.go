package main

import "fmt"

// 定义结构体
type Cat struct {
	Name  string
	Color string
	Age   uint
}

// new 构造结构体
func NewCat(name, color string, age uint) *Cat {
	// 传三个参数 构造了一个cat对象
	return &Cat{name, color, age}
}

func (c *Cat) Sleeping() {
	fmt.Println(c.Color, "Cat", c.Name, "is Sleeping")
}

func (c *Cat) Eating() {
	fmt.Println(c.Color, "Cat", c.Name, "is Eating")
}

func (c *Cat) Print() {
	fmt.Printf("+%v\n", c)
}
