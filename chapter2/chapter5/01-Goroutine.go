package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("l am main goroutine")
	// 一个 goroutine
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("l am a goroutine")
	}()
	fmt.Println("bye bye")
	// 睡眠确保其他的goroutin运行完成
	time.Sleep(time.Second * 2)
}
