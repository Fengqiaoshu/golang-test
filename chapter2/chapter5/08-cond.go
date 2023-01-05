package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 要做一个临界区：环形队列[0 1 2 3 4]
var five [5]int // 共享数据 - 饼框
var cond *sync.Cond
var mutex sync.Mutex

// 控制生产者生产数量
var cond_five *sync.Cond  // 条件变量
var mutex_five sync.Mutex //锁
var prod_count int = 0
var customer_index int = 0

// 生产者WaitGroup
func productor() {
	index := 0      // 生产下标
	prodnum := 1000 //产品标号从1000开始
	for {
		//抢锁
		mutex.Lock()
		//生产商品
		time.Sleep(time.Millisecond * 100)
		five[index] = prodnum
		fmt.Println(" l am productor,num====", prodnum)
		prodnum++
		index = (index + 1) % 5 // 环形队列
		cond.Signal()
		//释放锁
		mutex.Unlock()
		// 判断是否可以生产
		if prod_count == 5 {
			mutex_five.Lock()
			cond_five.Wait() // 生产者等待消费者通知
			mutex_five.Unlock()
		}
	}
	wg.Done()
}

// 消费者
func customer(num int) {

	for {
		// 抢锁
		mutex.Lock()
		// wait
		cond.Wait()
		if prod_count > 0 { // 防止出现极端情况
			// 吃饼
			time.Sleep(time.Millisecond * 10)
			fmt.Printf(" l am %d customer,num====%d\n", num, five[customer_index])
			customer_index = (customer_index + 1) % 5
			prod_count--
		}
		//释放锁
		mutex.Unlock()
		cond_five.Signal() //通知生产者

	}
	wg.Done()
}

func main() {
	cond = sync.NewCond(&mutex) //构造条件变量
	cond_five = sync.NewCond(&mutex_five)
	wg.Add(3)
	go productor()
	go customer(1)
	go customer(2)
	wg.Wait()
}
