package main

import "fmt"

//一个接口可以被多个结构体实现，一个类型可以实现多个接口
type Usb interface {
	read()
	write()
}

type computer struct {
	name string
}

func (c computer) read() {
	fmt.Printf("%v read...\n", c.name)
}
func (c computer) write() {
	fmt.Printf("%v write...\n", c.name)
}

func main6() {
	c := computer{
		name: "thinkpad",
	}
	c.read()
	c.write()
}
