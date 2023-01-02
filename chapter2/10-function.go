package main

import (
	"fmt"
)

func main() {
	x, y := 10, 20
	sum, sub := add_sub(x, y)
	fmt.Println(sum, sub)

	// 函数指针
	funcptr := add_sub
	sum1, sub1 := funcptr(x, y)
	fmt.Println(sum1, sub1)
	// funcptr 可以变为参数
	fmt.Println(count(100, 10, add_func))
	fmt.Println(count(100, 10, sub_func))
}

func add_sub(a, b int) (int, int) {
	return a + b, a - b
}

// 函数 内部：值，并且如何计算
func count(a, b int, math func(a, b int) int) int {
	return math(a, b)
}

// 遵循函数 参数的要求，函数要作为参数
func add_func(a, b int) int {
	return a + b
}

func sub_func(a, b int) int {
	return a - b
}
