package main

import "fmt"


func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("出错")
	fmt.Println(4)
}
