// 完成生产者消费者过程
package semaphore

import (
	"fmt"
	"time"
)

func Producer(id int, item chan int) {
	for i:=0; i<10; i++ {
		item <- i
		fmt.Printf("Producer %d produces data: %d\n", id, i)
		time.Sleep(1 * time.Second)
	}
}

func Consumer(id int, item chan int) {
	for i:=0; i<20; i++ {
		cItem := <-item
		fmt.Printf("Consumer %d get data: %d\n", id, cItem)
		time.Sleep(1 * time.Second)
	}
}

func Start() {
	channel := make(chan int, 5)
	go Producer(111, channel)
	go Producer(222, channel)
	go Consumer(333, channel)
	time.Sleep(20 * time.Second)
}