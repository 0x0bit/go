package main

import (
	"fmt"
	"math"
)

// 常量定义在函数内部
func consts() {
	const fileName = "abc.txt"
	const a, b = 3, 4
	var c int

	// 注：这里因为const 声明a,b的时候并未指明数据类型，所以如下的时候可以不用强制数据类型转换
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println("函数内部常量", fileName, c) //函数内部常量 abc.txt 5
}

//常量定义在函数外部
const fileName = "abcd.txt"
const a, b = 3, 4

func consts2() {
	c := int(math.Sqrt(a*a + b*b))
	fmt.Println("函数外部常量", fileName, c) // 函数外部常量 abcd.txt 5
}

// 也可以简写
const (
	fileName2 = "abc.txt"
	d, f      = 3, 4
)

func consts3() {
	c := int(math.Sqrt(d*d + f*f))
	fmt.Println("简写外部常量", fileName2, c) // 简写外部常量 abc.txt 5
}

// -------------------------------------------------------------
// 枚举
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang) //0 1 2 3
}

// 对于连续的常量，go提供了一种简洁的写法
func enums2() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang) //0 1 2 3
}

// 跳过一个值
func enums3() {
	const (
		cpp = iota
		_
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang) //0 2 3 4
}

// iota 也可以作为一个自增值的种子
func enums4() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb) //1 1024 1048576 1073741824 1099511627776 1125899906842624
}

func main() {
	consts()
	consts2()
	consts3()
	enums()
	enums2()
	enums3()
	enums4()
}
