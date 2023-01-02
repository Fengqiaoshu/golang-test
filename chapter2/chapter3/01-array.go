package main

import (
	"fmt"
)

func main() {
	// 定义五个元素的整形数组
	var five [5]int = [5]int{1, 2, 3, 4, 5}
	// 定义三个元素的字符串数组
	var three [3]string = [3]string{"lucy", "lily", "lilei"}
	var two [3]string = [3]string{"lucy", "lily"}
	fmt.Println(five)
	fmt.Println(three)
	fmt.Println(two)
}
