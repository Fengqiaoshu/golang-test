package main

import (
	"fmt"
)

// 定义动物接口

type Animal interface {
	sleeping()
	eating()
}
type Dog struct {
	color string
}
type Cat struct {
	color string
}

// dog支持接口 方法去支持接口
func (d Dog) sleeping() {
	fmt.Printf("%s's dog is sleeping\n", d.color)
}
func (d Dog) eating() {
	fmt.Printf("%s's dog is eating\n", d.color)
}
func (c Cat) sleeping() {
	fmt.Printf("%s's cat is sleeping\n", c.color)
}
func (c Cat) eating() {
	fmt.Printf("%s's cat is eating\n", c.color)
}

// 工厂模式：流水线作业，可以生产猫 可以生产狗 返回值怎么办

func factory(color string, animal_type string) Animal {
	switch animal_type {
	case "dog":
		return &Dog{color}
	case "cat":
		return &Cat{color}
	default:
		return nil // nil是空值
	}
}

func main() {
	d1 := Dog{
		color: "green",
	}
	d1.eating()
	c1 := Cat{
		color: "red",
	}
	c1.sleeping()
	fmt.Println("--------------")
	c2 := factory("yellow", "cat")
	c2.eating()
	d2 := factory("black", "dog")
	d2.sleeping()
}
