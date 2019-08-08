package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["name"] = "Tom"
	m["sex"] = "M"
	m["hbs"] = "打球"

	fmt.Println("m= ", m)

	m2 := map[string]string {
		"name":"tom",
		"site":"imooc",
	}
	fmt.Println("m2= ", m2)

	// 遍历map
	for k, v := range m {
		fmt.Println("k =", k, "v = ", v)
	}

	name := m["name"]
	fmt.Println(name)

	// 判断这个键存不存在
	age, ok := m2["age"]
	fmt.Println("age, ok =", age, ok)

	delete(m2, "site")
	fmt.Println(m2)
}
