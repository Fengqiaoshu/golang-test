package main

import (
	"fmt"
	"os"
	"time"
)

// 监控标准输入
func cancel(out chan<- string) {
	buf := make([]byte, 10)
	os.Stdin.Read(buf) //阻塞读
	//通道通知主控的goroutine
	out <- "stop"

}

func main() {
	stdin_chan := make(chan string)
	go cancel(stdin_chan)
	ticker := time.NewTicker(time.Second) // 每隔1s的定时器
	num := 5
	for num > 0 {
		//data := <-ticker.C
		// 可以监控多路channel
		select {
		case <-ticker.C:
			fmt.Println(num)
			num--
		case <-stdin_chan:
			ticker.Stop()
			return
		}
	}
	fmt.Println("发射!!!!")
	ticker.Stop()

}
