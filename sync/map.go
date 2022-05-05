package main

import (
	"fmt"
	"sync"
)

func main() {
	 syncMap := sync.Map{}

	 // Load(): 通过key获取value
	 // 获取一个不存在的key的值
	 v, ok := syncMap.Load("a")
	 fmt.Printf("Load() v=%v, bool=%v\n", v, ok)

	 // Store(): 设置key-val
	 syncMap.Store("a", 1)
	 syncMap.Store("b", 2)
	 v, ok = syncMap.Load("a")
	 fmt.Printf("Load() v=%v, bool=%v\n", v, ok)

	 // Range()：便利map中的key-val
	 syncMap.Range(func(key, value interface{}) bool{
		fmt.Printf("Range() k=%v, v=%v\n", key, value)
		return true
	 })

	 // LoadOrStore(): 设置并存储
	 // 设置一个存在的key, key存在，不进行重置操作，值还是之前的1，b=true
	 act, b := syncMap.LoadOrStore("a", 2)
	 fmt.Printf("LoadOrStore v=%v, bool=%v\n", act, b)

	 // 设置一个不存在的key, 值还是2，b=false
	act, b = syncMap.LoadOrStore("c", 2)
	fmt.Printf("LoadOrStore v=%v, bool=%v\n", act, b)

	 // Delete() 删除key
	 syncMap.Delete("a")
}