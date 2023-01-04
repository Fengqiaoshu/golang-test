package main

import (
	"fmt"
	"sync"
)

// 命名全局变量
var x int = 0
var wg sync.WaitGroup

// 定义锁
var mutex sync.Mutex

/*
	加锁步骤：
		1. 请求修改前加锁
		2. 修改数据	(加锁，临界区一定要最小颗粒度)
		3. 解锁
*/

func increment() {
	mutex.Lock()
	x += 1
	mutex.Unlock()
	wg.Done()

}

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait() //确保运行完成
	fmt.Println("final value is", x)
}
