package main

import (
	"fmt"
)
func main(){
	// 常量的声明方式
	const 	LENGTH int = 10
	const	WIDTH = 20
	const a,b,c = 3.14,false,"hello"

	// 
	area := LENGTH * WIDTH
	fmt.Println(area)
	fmt.Println(a,b,c)

}