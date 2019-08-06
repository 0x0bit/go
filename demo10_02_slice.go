package main

import "fmt"

/**
对slice的一些基本操作，如创建、增加、删除
 */
func main() {
	// 创建方式1
	var s1  []int
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
	}
	fmt.Println("s1=",s1, "len=", len(s1), "cap=", cap(s1))

	// 创建方式2
	s2 := []int{2, 3, 4, 5, 6}
	fmt.Println("s2=",s2, "len=", len(s2), "cap=", cap(s2))

	// 创建方式3，使用make
	s3 := make([]int, 6, 10)
	fmt.Println("s3=",s3, "len=", len(s3), "cap=", cap(s3))

	// copy
	copy(s3, s2)
	fmt.Println(s2)
}

/**
运行结果：
	s1= [0 1 2 3 4 5 6 7 8 9] len= 10 cap= 16
	s2= [2 3 4 5 6] len= 5 cap= 5
	s3= [0 0 0 0 0 0] len= 6 cap= 10
	[2 3 4 5 6]

*/
