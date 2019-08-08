package main

import (
	"fmt"
	"strconv"
)
// 动物类
type Animal struct {
	name string
	subject string
}

// 动物的公共方法
func (a *Animal) eat(food string) {
	fmt.Println(a.name + "喜欢吃：" + food +",它属于:" + a.subject)
}

// 猫类，继承动物类
type Cat struct {
	Animal
	age int
	ss func()
}

// 猫类独有的方法
func (c Cat) sleep() {
	fmt.Println(c.name + " 今年" + strconv.Itoa(c.age) + "岁了,特别喜欢睡觉")
}

func main() {
	// 创建一个动物类
	animal := Animal{name:"动物", subject:"动物科"}
	animal.eat("肉")

	// 创建一个猫类
	cat := Cat{Animal: Animal{name:"咪咪", subject:"猫科"},age:1}
	cat.eat("鱼")
	cat.ss = func() {
		fmt.Println(111)
	}

	cat.sleep()

	/**
		动物喜欢吃：肉,它属于:动物科
		咪咪喜欢吃：鱼,它属于:猫科
		咪咪 今年1岁了,特别喜欢睡觉
	*/
}