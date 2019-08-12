package main

import "fmt"

func write1(ch chan int)  {
	for i := 0; i < 20; i++{
		ch <- i
	}
	close(ch)
}

func read1(ch chan int){
	for v := range ch{
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 1)
	go write1(ch)
	read1(ch)
}
