package main // package

import (
	"fmt"
	"os"
) // dependency

// main function
func main()  {
	if len(os.Args)>1 {
		fmt.Println("Hello World", os.Args[1])
	}
	os.Exit(0)
}

//运行
//$ go run hello_world.go
//Hello World
//编译
//go build hello_world.go
//$ ls
//hello_world.exe*  hello_world.go

//注意: 程序入口必须遵循 : 1. main package 2. main function
//差异: 1. main方法无返回值 2. 使用os.Exit() 返回状态 3. 使用os.Args传递命令行参数
