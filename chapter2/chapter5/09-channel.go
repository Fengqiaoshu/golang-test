package main

import (
	"fmt"
	"time"
)

// 定义一个通道
var c chan string

// 读通道
func reader() {
	fmt.Println("begin read")
	msg := <-c
	fmt.Println("read:", msg)
}

func main() {
	c = make(chan string)       // 创建无缓冲通道
	go reader()                 // 启动goroutine
	time.Sleep(time.Second * 1) // 为了演示效果
	c <- "hello"                // 写入通道	- 有可能阻塞
	time.Sleep(time.Second * 3)
}
