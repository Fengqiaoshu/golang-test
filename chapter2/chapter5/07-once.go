package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int = 0 // 共享数据
	var wg sync.WaitGroup
	var once sync.Once

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			once.Do(func() {
				count += 1
			})
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count=", count)
}
