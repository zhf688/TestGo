package main

import (
	"fmt"
	"os"
	"strconv"
)

func main2() {
	var score int = 0
	var fenshu string = "A"
	fmt.Println("欢迎进入成绩查询系统===============")

ZIZHU:
	for true {
		var xuanze int = 0
		fmt.Println("1、进入成绩查询,2、退出程序")
		fmt.Println("请选择：")
		fmt.Scanln(&xuanze)
		if xuanze == 1 {
			goto CHA
		} else {
			os.Exit(1)
		}
	}

CHA:
	for true {
		fmt.Println("请输入要查询的成绩：")
		fmt.Scanln(&score)
		switch {
		case score >= 90:
			fenshu = "A"
		case score >= 80 && score < 90:
			fenshu = "B"
		case score >= 60 && score <= 80:
			fenshu = "C"
		default:
			fenshu = "D"
		}

		var c string = strconv.Itoa(score)
		switch {
		case fenshu == "A":
			fmt.Printf("你考了%s分，评价为%s，成绩优秀\n", c, fenshu) //格式化输出
		case fenshu == "B":
			fmt.Printf("你考了%s分，评价为%s，成绩良好\n", c, fenshu)
		case fenshu == "C":
			fmt.Printf("你考了%s分，评价为%s，成绩合格\n", c, fenshu)
		case fenshu == "D":
			fmt.Printf("你考了%s分，评价为%s，成绩不合格\n", c, fenshu)
		}
		goto ZIZHU
	}
}
