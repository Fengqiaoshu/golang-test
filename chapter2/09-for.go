package main

import (
	"fmt"
	"time"
)

func main() {

	i, sum := 0, 0
	// 第一种
	for i = 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	// 第二种
	i = 1
	sum = 0
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
	// 每个1s 打印一个haha
	for {
		fmt.Println("haha")
		// 睡眠1s
		time.Sleep(time.Second * 1)

	}
}
