package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(time.Millisecond * 100) // 启动一个观感服务
	fmt.Println("\n", fib(45))         // 时间长，影响体验
}

func fib(x int) int {
	// f(x) = f(x-1) + f(x-2) x=0 f=0 x=1 f=1
	if x < 2 {
		return x
	}

	return fib(x-2) + fib(x-1)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
