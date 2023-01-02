package main

import "fmt"

func main() {
	a1 := [5]int{1, 2, 3, 4, 5} //定义一个数组
	// 切片
	s1 := a1[2:4] // s1 == {3,4} = [start_index:end_index - 1]
	// = [start_index:start_index + length] = [2:4] => [2:2+2]

	fmt.Println(s1)
	fmt.Println("---------------")
	s1[1] = 101

	fmt.Println(s1)
	fmt.Println(a1)

}
