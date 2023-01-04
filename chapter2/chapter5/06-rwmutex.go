package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var rwlock sync.RWMutex
var x int = 0 // 共享变量

func main() {
	wg.Add(10) // 监控10个goroutine
	// 启动 7个读的 goroutine
	for i := 0; i < 7; i++ {
		go reader(i)
	}
	// 启动三个写的 goroutine
	for i := 0; i < 3; i++ {
		go writer(i)
	}
	wg.Wait()
}

func reader(num int) {
	for {
		rwlock.RLock() // 读锁
		fmt.Printf("l am %d reader goroutine x = %d\n ", num, x)
		time.Sleep(time.Millisecond * 2)
		rwlock.RUnlock() // 释放读锁
	}
	wg.Done()
}

func writer(num int) {
	for {
		rwlock.Lock() // 写锁
		x += 1
		fmt.Printf("l am %d writer goroutine x = %d\n", num, x)
		time.Sleep(time.Millisecond * 2)
		rwlock.Unlock() // 释放读锁
	}
	wg.Done()
}
