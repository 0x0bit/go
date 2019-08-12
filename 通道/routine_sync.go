package main

import (
	"fmt"
	"sync"
	"time"
)

func write3(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()	// 计数值减1
	i := 0
	for i < 4 {
		i++
		ch <- i
	}
}

func read3(ch chan int) {
	for v := range ch{
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 4)
	wg := new(sync.WaitGroup)
	wg.Add(2) // 增加计数值
	// 增加两个协程
	go write3(ch, wg)	// 写
	go write3(ch, wg)	// 写

	go read3(ch)
	// wait() 阻塞等待所有的写通道协程结束，等待数值变为0，Wait()才会返回
	wg.Wait()
	//关闭通道
	close(ch)
	time.Sleep(time.Second)
}