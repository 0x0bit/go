package main

import "fmt"

func write(ch chan int) {
	i := 0
	for {
		i++
		ch <- i
	}
}

func read(ch chan  int) {
	value := <- ch
	fmt.Println(value)
	close(ch)
}

func main(){
	var ch  = make(chan int, 4)
	go read(ch)
	write(ch)
}
