
package main

import (
	"fmt"
	"os"
)

func main() {

	// 获取hostname
	fmt.Println(os.Hostname())   // localhost <nil>

	// 返回底层系统内存页的大小
	fmt.Println(os.Getpagesize())  // 4096

	// 返回环境变量, key=val的格式的切片
	fmt.Println(os.Environ())

	// 获取某个环境变量的值
	fmt.Println(os.Setenv("demo", "demo"))

	// 设置某个环境变量的值，如果之前的key已存在，则会覆盖
	fmt.Println(os.Getenv("demo"))

	// 程序退出，以给定状态码的形式，让程序退出。exit后面的代码不执行，包括defer
	// os.Exit(1)

	// 返回调用者的用户id
	fmt.Println(os.Getuid())    // 501

	// 返回调用者的有效的用户id
	fmt.Println(os.Geteuid())   // 501

	// 返回调用者的组id
	fmt.Println(os.Getgid())   // 20

	// 返回调用者的有效组id
	fmt.Println(os.Geteuid())  // 501

	// 返回调用者所在的进程id
	fmt.Println(os.Getpid())

	// 返回调用者所在进程的父进程id
	fmt.Println(os.Getppid())

	// FileInfo相关操作
	filePath := "os.go"
	fileInfo, err  := os.Stat(filePath)

	// 获取文件名
	fmt.Println(fileInfo.Name())   // os.go

	// 返回文件的模式位，返回FileMode类型
	fmt.Println(fileInfo.Mode())   // -rw-r--r--

	// 文件的修改时间
	fmt.Println(fileInfo.ModTime())  // 2021-01-08 18:24:50.416467241 +0800 CST

	// 是否是一个目录
	fmt.Println(fileInfo.IsDir())   //  false

	// 是否是一个目录
	fmt.Println(fileInfo.Mode().IsDir())  // false

	// 根据错误判读文件/目录是否存在
	fmt.Println(os.IsExist(err), os.IsNotExist(err)) // false  false

	// 根据错误判断是否是因为权限导致
	fmt.Println(os.IsPermission(err))   // false

	// 返回当前文件的路径
	fmt.Println(os.Getwd())
}