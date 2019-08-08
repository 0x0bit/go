package main

import "fmt"

type USB interface {
	start()
	end()
}

// 实现
// 鼠标
type Mouse struct {
	name string
}

// U盘
type FlashDisk struct {
	name string
}

// 鼠标按照接口生产
func (m Mouse) start() {
	fmt.Println(m.name, "鼠标，准备就绪，可以开始工作，可以开始点点点。。。。。。")
}
func (m Mouse) end() {
	fmt.Println(m.name, "结束工作，可以安全退出。。")
}

// U盘按照接口生产
func (f FlashDisk) start() {
	fmt.Println(f.name,"准备开始工作，可以进行数据存储了。。")
}

func (f FlashDisk) end() {
	fmt.Println(f.name,"可以弹出。。。")
}

// 多态
func testInterface(usb USB) {
	usb.start()
	usb.end()
}

func main() {
	mouse := Mouse{name:"雷蛇鼠标"}
	testInterface(mouse)

	uPan := FlashDisk{name:"闪迪U盘"}
	testInterface(uPan)
}