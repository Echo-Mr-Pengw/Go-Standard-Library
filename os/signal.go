package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	c := make(chan os.Signal)

	// Notify使用
	//// 监听信号
	//// Notify如果只有第一个参数，表示监听所有的中断；否则，监听指定的中断
	//signal.Notify(c, os.Interrupt, os.Kill)
	//
	//fmt.Println("开始接受信号")
	//// 未收到中断前，阻塞
	//signal := <-c
	//fmt.Println("信号为：", signal)


	// Notify、Stop的使用

	signal.Notify(c)
	// 停止像c中转发信号
	signal.Stop(c)
	// 此处会阻塞，没有信号
	signal := <-c

	fmt.Println("信号为：", signal)
}