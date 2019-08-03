package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	c := 3 + 4i
	// 复数求模
	fmt.Println(cmplx.Abs(c)) // 5

	// 欧拉公式
	b := cmplx.Pow(math.E, 1i*math.Pi) + 1
	// 等价于
	d := cmplx.Exp(1i*math.Pi) + 1

	fmt.Println(cmplx.Abs(b))
	fmt.Println(cmplx.Abs(d)) // 1.2246467991473515e-16
	fmt.Printf("%.3f\n", d)   //(0.000+0.000i)
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c) // 5
}

func main() {
	euler()
	triangle()
}
