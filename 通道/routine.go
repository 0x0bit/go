package main

import "fmt"

func main (){
	var ch chan int = make(chan int, 4)
	for i:=0; i<cap(ch); i++{
		ch <- i		// 写通道
		fmt.Println("写", i)
	}
	fmt.Println("ch len",len(ch))
	for len(ch) > 0{
		var value int = <-ch
		fmt.Println("读", value)
	}
	fmt.Println("ch len", len(ch))
}
