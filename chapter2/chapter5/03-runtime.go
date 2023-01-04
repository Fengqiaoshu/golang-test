package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("max proc is", runtime.GOMAXPROCS(1))
	runtime.GOMAXPROCS(3) // 设置 小于等于1的时候是查看  大于1的时候是设置
	fmt.Println("max proc is", runtime.GOMAXPROCS(1))
	//runtime.Goexit()	退出 goroutine
	runtime.Gosched() // 释放cpu
}
