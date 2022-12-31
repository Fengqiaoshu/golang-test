package main

import (
	"fmt"
)

func main() {
	var v1 uint
	var x,y = 123,"hello"
	u,t := 456,"world"
	fmt.Println(v1,x,y,u,t)
	var c1 complex64 = 3 +4i  // 4i要连着
	fmt.Println(c1)
}
