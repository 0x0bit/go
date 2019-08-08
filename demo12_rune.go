package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	s := "Yes我爱大中国"
	for _, b := range []byte(s){
		fmt.Printf("%X ", b)
		// 59 65 73 E6 88 91 E7 88 B1 E5 A4 A7 E4 B8 AD E5 9B BD
	}
	fmt.Println()
	fmt.Println(utf8.RuneCountInString(s))
	for i, ch := range []rune(s){
		fmt.Printf("(%d, %c)", i, ch)
	}
	fmt.Println()
}
