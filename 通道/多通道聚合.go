package main

import (
	"fmt"
	"time"
)

// 每隔一会生产一个数
func write4(ch chan int, gap time.Duration){
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(gap)
	}
}

// 将多个原通道内容拷贝到单一的目标通道
func collect(source chan int, target chan int){
	for v := range source{
		target <- v
	}
}

// 从目标通道消费
func read4(ch chan int){
	for v := range ch{
		fmt.Printf("receive %d\n", v)
	}
}

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go write4(ch1, time.Second)
	go write4(ch2, 2 * time.Second)

	go collect(ch1, ch3)
	go collect(ch1, ch3)

	read4(ch3)
}