// 日志相关操作
// 日志的三类接口：Print、Fatal、Panic
// 每类接口对应三种方法：Print Println Printf、Fatal Fatalf Fatalln  Panic Panicf Panicln
// Print类：跟fmt差不多，只是前面加了格式
// Fatal类：日志输出后，系统调用os.exit(1)，整个程序退出。如果后面有defer,也不执行
// Panic类：日志输出后，发生Panic

package main

import (
	"log"
	"os"
)

func main() {
	// 默认的日志输出
	//log.Print("Print 日志")      // 2020/11/24 17:27:26 Print 日志
	//log.Println("Println 日志")  // 2020/11/24 17:27:26 Println 日志
	//log.Printf("Printf 日志")    // 2020/11/24 17:27:26 Printf 日志
	//
	//log.Fatal("Fatal 日志")              // 2020/11/24 17:30:35 Fatal 日志  exit status 1
	//log.Printf("Fatal执行后，是否显示？")  // 不执行
	//
	//log.Panic("Panic 日志")
	//
	//// 设置输出的日志前缀
	//log.SetPrefix("go:")
	//log.Println("日志前缀")  // go:2020/11/24 18:22:54 日志前缀
	//// 打印日志的前缀
	//fmt.Println(log.Prefix())
	//// 标志位
	//fmt.Println(log.Flags())

	// 自定义日志输出
	logger := log.New(os.Stdout, "log:", log.Lshortfile)  // log:log.go:37: 自定义日志
	logger.Println("自定义日志")
	// 查看自定义的日志前缀
	logger.Println(logger.Prefix())  // log:log.go:39: log:
}

