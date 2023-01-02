package main

import (
	"fmt"
)

func main() {
	a, b := 10, 20
	a, b = swap(a, b)
	fmt.Println(a, b)
}

// Go语言多个返回值，多个接收者
func swap(x, y int) (int, int) {
	return y, x
}
