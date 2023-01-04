package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(num int) {
			time.Sleep(time.Second * time.Duration(num))
			fmt.Println("l am", num, "goroutine!")
			wg.Done() // --
		}(i)
	}
	wg.Wait() // 阻塞等待所有goroutine退出
	fmt.Println("-------------")
}
