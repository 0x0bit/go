package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile (fileName string) {
	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	// 最后关闭文件
	defer file.Close()

	// 写文件
	write := bufio.NewWriter(file)

	for i := 0; i < 20; i++ {
		fmt.Fprintln(write, i)
	}
	//写完之后刷新
	defer write.Flush()
}

func main()  {
	writeFile("abc.txt")
}