package main

/*
	笔记地址： https://www.yuque.com/continue/go/in99xc
	数组的切片
*/
import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]:", arr[2:6]) // 从下标2到下标6
	fmt.Println("arr[:6]:", arr[:6])   // 从下标0到下标6
	fmt.Println("arr[2:]:", arr[2:])   // 从下标2到最后
	fmt.Println("arr[:]:", arr[:])     // 全部

	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("s1--->", s1) // s1---> [2 3 4 5]
	fmt.Println(len(s1), cap(s1))
	fmt.Println("s2--->", s2) // s2---> [5 6]

}
