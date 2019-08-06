package main

/**
笔记： https://www.yuque.com/continue/go/tv4iv3
流程控制结构
	- if
	- if……else
	- switch
		- switch...case： 后不需要 break
*/
import (
	"fmt"
	"io/ioutil"
)

func ifTest1() {
	const fileName = "abc.txt"
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("%s", contents)
	}
}

// 写法2， 这种写法如果出了 if 则无法在访问
func ifTest2() {
	const fileName = "abc.txt"
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("%s", contents)
	}
}

// switch,会自动break
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("未知符号: " + op)
	}
	return result
}

// switch 后可以没有表达式
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("worong score: %d", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	ifTest1()
	ifTest2()

	fmt.Println(
		grade(0),
		grade(69),
		grade(100),
		grade(123),
	)
}
