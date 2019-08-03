package main

import "fmt"

var (
	aa = 12
	bb = "hello world"
)

// 声明变量不附初始值
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 声明变量附初始值
func variableInitValue() {
	var a int = 3
	var s string = "abc"
	fmt.Println(a, s)
}

// 声明变量不指定变量类型
func variableTypeDeduction() {
	var a, b, c, d = 2, 2.3, true, "def"
	fmt.Println(a, b, c, d)
}

// 省略关键字 var 声明变量，但是必须在函数体中使用这种形式来声明变量，全局变量不能用这种形式来声明。
func variableShorter() {
	a, b, c, d := 2, 2.3, true, "def"
	b = 5
	fmt.Println(a, b, c, d)
}

func main() {
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb)
}

/* 输出结果
0 ""
3 abc
2 2.3 true def
2 5 true def
12 hello world
*/
