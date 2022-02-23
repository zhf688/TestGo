package main

import (
	"fmt"
	"os"
)

func main() {
	//获取所有环境变量
	fmt.Printf("os.environ:%v\n", os.Environ())
	//根据环境变量名称查找
	fmt.Printf("JAVA_HOME:%v\n", os.Getenv("JAVA_HOME"))
	//lookupenv查找，有结果才打印
	env, b := os.LookupEnv("JAVA_HOME")
	if b {
		fmt.Printf("env:%v\n", env)
	}
	//设置环境变量
	os.Setenv("ENV1", "env1...")
	os.Setenv("ENV2", "env2...")
	//os.expandenv扩展环境变量
	fmt.Printf(os.ExpandEnv("$ENV1 links $ENV2"))

	//os.Clearenv() 清空环境变量
}
