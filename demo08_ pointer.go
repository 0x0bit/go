package main

// 指针
// 笔记： https://www.yuque.com/continue/go/zrflqm

import (
	"fmt"
)

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a, pa)
}
