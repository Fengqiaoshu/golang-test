// 闭包
package main

import (
	"fmt"
)

func main() {
	// 获得函数指针
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	fmt.Println("---------------")
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

func getSequence() func() int {
	i := 0 //函数内部变量

	return func() int {
		i += 1
		return i
	}
}
