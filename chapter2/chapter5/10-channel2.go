package main

import (
	"fmt"
	"time"
)

var c1 chan int
var c2 chan int

func main() {
	//通道初始化
	c1 = make(chan int)
	c2 = make(chan int)
	//启动goroutine
	// 通道是可以关闭的 只能在写端关闭
	//  counter
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i //写入
			time.Sleep(time.Second)
		}
		// 关闭通道
		close(c1)
	}()
	// squarer
	go func() {
		for {
			num, ok := <-c1 // 读取可以设置标志位
			if !ok {
				break //如果通道异常
			}
			c2 <- num * num
		}
		close(c2)
	}()
	// printer
	for {
		num, ok := <-c2
		if !ok {
			break //如果通道异常
		}
		fmt.Println(num)
	}

}
