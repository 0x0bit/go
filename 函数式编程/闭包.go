package main

import "fmt"

/*
对于函数adder() ,他的返回值是一个函数，这个函数接收一个int值，并返回一个int值。
函数内部的第一个return 是对应adder(),因为它需要返回一个函数
第二个return 对应的是 返回函数，因为返回函数需要返回一个int值
 */
func adder() func(value int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

func main() {
	adder := adder()
	for i := 0; i< 10; i++ {
		fmt.Println(adder(i))
	}
}