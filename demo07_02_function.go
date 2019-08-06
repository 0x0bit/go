package main

import (
	"fmt"
)

func evalNum2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div2(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("未知的运算符：%s", op)
	}
}

// 求商和余数
func div2(a, b int) (q, r int) {
	return a / b, a % b
}

// 可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := evalNum2(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(div2(13, 3))
}
