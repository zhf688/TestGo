package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("os.getpid:%v\n", os.Getpid())
	fmt.Printf("os.getppid:%v\n", os.Getppid())
	//设置新打开进程的属性
	attr := &os.ProcAttr{
		//files指定新进程继承的活动文件对象
		//标准输入，标准输出，标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Env:   os.Environ(),
	}

	//开启一个新的进程
	process, err := os.StartProcess("E:/Program Files/Notepad++/notepad++.exe",
		[]string{"E:/Program Files/Notepad++/notepad++.exe", "D:/a.txt"}, attr)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("p进程id:%v\n", process.Pid)

	//通过进程号查找进程
	findProcess, _ := os.FindProcess(process.Pid)
	fmt.Printf("findprocess:%v\n", findProcess)

	//等待10秒后执行函数
	time.AfterFunc(time.Second*10, func() {
		//向进程发送信号
		process.Signal(os.Kill)
	})
	wait, _ := process.Wait() //等待进程结束后，返回状态
	fmt.Printf("wait:%v\n", wait)
}
