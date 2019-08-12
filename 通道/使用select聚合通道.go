package main

import (
	"fmt"
	"time"
)

func write5(ch chan int, gap time.Duration){
	i := 0
	for{
		i++
		ch <- i
		time.Sleep(gap)
	}
}

func read5(ch1 chan int, ch2 chan int){
	for{
		select {
		case v := <- ch1:
			fmt.Printf("read %d from ch1\n", v)
		case v := <- ch2:
			fmt.Printf("read %d from ch2\n", v)
		default:

		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go write5(ch1, time.Second)
	go write5(ch2, time.Second)

	read5(ch1, ch2)
}
