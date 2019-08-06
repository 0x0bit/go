package main

import (
	"fmt"
)

func updateArr(arr [5]int) {
	arr[0] = 100
	fmt.Println("修改后的arr：", arr)
}

func main() {
	arr3 := [...]int{2, 4, 5, 6, 7}

	// 使用range关键字同时回去下标和值
	for i, v := range arr3 {
		fmt.Println(i, v)
		/*
			结果：
					0 2
					1 4
					2 5
					3 6
					4 7
		*/
	}

	fmt.Println("原来的：", arr3)
	updateArr(arr3)
	fmt.Println("再次查看原始的：", arr3)
}
