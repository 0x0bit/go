package main

/*
笔记：https://www.yuque.com/continue/go/zb108r
	- for
		- 可以省略起始条件
		- 可以省略自增

	- 没有while 循环
		- for 循环只有条件语句的时候
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 十进制转二进制， for循环使用起始条件
func convertToBin1(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		// 字符串转换
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// for 循环省略递增条件
func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin1(5),  // 101
		convertToBin1(13), // 1101
		convertToBin1(0),  // 空串
	)

	printFile("abc.txt")
}
