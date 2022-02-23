package main

import "os"

//打开文件并写入字节数组
func writeByte() {
	//os.O_APPEND表示在末尾追加， os.O_TRUNC表示清空
	file, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755)
	file.Write([]byte(" hello aaa"))
	file.Close()
}

//写入字符串
func writeString() {
	file, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755)
	file.WriteString("===hello bbb")
	file.Close()
}

//在指定位置写入
func writeAt() {
	//使用WriteAt方法，不能用os.O_APPEND
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0755)
	f.WriteAt([]byte("..."), 3)
	f.Close()
}
func main() {
	//writeByte()
	//writeString()
	writeAt()
}
