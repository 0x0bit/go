package main

import "fmt"

/**
为结构定义方法
 */

type person struct {
	name string
	age int
	gender string
}

// 方法定义
func (p *person) describe() {
	fmt.Printf("我叫 %v, 我今年 %v 了", p.name, p.age)
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p *person) setName(name string) {
	p.name = name
	fmt.Println("person name is:", p.name)
}

func main() {
	p := person{"tom", 22, "man"}
	p.setAge(25)
	p.setName("jury")
	p.describe()
}
