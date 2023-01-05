package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second) // 每隔1s的定时器
	num := 5
	for num > 0 {
		//data := <-ticker.C
		<-ticker.C
		fmt.Println(num)
		num--
	}
	fmt.Println("发射!!!!")
	ticker.Stop()

}
