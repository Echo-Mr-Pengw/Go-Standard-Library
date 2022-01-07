package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var addr int32 = 21

	// 原子的获取addr的值
	v1 := atomic.LoadInt32(&addr)
	fmt.Println(v1)

	// 原子的将值val保存到addr
	var v2 int32
	atomic.StoreInt32(&v2, addr)
	fmt.Println(v2, addr)

	// 给addr加10
	v3 := atomic.AddInt32(&addr, 10)
	fmt.Println(v3)
}