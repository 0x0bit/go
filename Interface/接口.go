package main

import "fmt"

type Person interface {
	eat()
	sleep()
}

type Student struct {}

func (stu Student) eat(){
	fmt.Println("学生是人，也要吃饭")
}

func (stu Student) sleep()  {
	fmt.Println("学生是人，也要睡觉")
}

func main() {
	stu := Student{name: "哈哈"}
	stu.eat()
	stu.sleep()
}