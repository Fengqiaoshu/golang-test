package main

import (
	"fmt"
)

func main() {
	var fruit string
	// 标准输入
	//os.Stdin
	fmt.Print("plese input a fruit's name")
	fmt.Scanf("%s", &fruit) //等待标准输入，赋值给fruit

	switch fruit {
	case "apple":
		fmt.Println("A,you choose apple")
	case "banana":
		fmt.Println("B,you choose banana")
	case "orange":
		fmt.Println("O,you choose orange")
	default:
		fmt.Println("Are you kidding me?")
	}
}
