package main

import "fmt"

// 定义结构体
type Person struct {
	name  string
	age   int
	sex   string
	fight int
}

func main() {
	p := Person{
		name:  "韩立",
		age:   25,
		sex:   "男",
		fight: 10000,
	}

	fmt.Println(p)

	p2 := Person{"lilei", 25, "man", 5}
	fmt.Printf("%+v\n", p2)
}
