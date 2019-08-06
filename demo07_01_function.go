package main

/*

笔记地址：https://www.yuque.com/continue/go/fnv7rk

	函数的基本用法
		- 返回一个值 和返回两个值的写法
*/

import (
	"fmt"
)

func evalNum(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("未知的运算符：" + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	fmt.Println(evalNum(3, 4, "*"))
	q, r := div(13, 3)

	fmt.Println(q, r)

	// 只接收一个返回值
	c, _ := div(13, 3)
	fmt.Println(c)

}
