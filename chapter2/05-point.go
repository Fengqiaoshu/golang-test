package main

import (
	"fmt"
)

func main() {
	// 指针
	a:=10
	// b指向a的地址
	var b *int = &a 
	fmt.Println(*b,b)
	*b = 100
	fmt.Println(a,*b,b)
}


