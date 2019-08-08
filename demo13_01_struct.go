package main

import "fmt"

type Persion struct {
	name string
	age int
	gender string
}

func main() {
	p1 := Persion{name:"tom", age:23, gender:"M"}
	p2 := Persion{"Jury", 23, "W"}

	fmt.Println(p1, p2)
	fmt.Println("name:", p1.name)

	p := &Persion{name:"tom", age:23, gender:"M"}
	fmt.Println(p.name)
}
