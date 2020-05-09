// RWMutex 读写锁
// RLock   读锁。有写操作需要等到读锁释放；有读操作不需要等待。
// RUnlock 释放读加锁。在未释放读锁之前，所有的写操作需要等待。
// Lock    写锁。加写锁后，不管是读还是写操作，必须等待写锁被释放。
// Unlock  释放写锁

// 此案例介绍写锁

package main

import (
	"fmt"
	"sync"
	"time"
)

var wmutex = sync.RWMutex{}
var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)

	go Room(1)
	go Room(2)

	wg.Wait()

	fmt.Println("main 执行完毕")
}

func Room(name int) {

	defer wg.Done()
	fmt.Println(name, "开始进去房间.....")
	// 加写锁
	wmutex.Lock()
	time.Sleep(time.Second)

	fmt.Println(name, "已经进去房间.....")
	fmt.Println(name, "已经离开房间")
	wmutex.Unlock()
}

// 通过下面的打印结果可以看出，2进入房间后加上写锁，如果还有其它协程需要进去，则需等到2出房间
// output:
// 2 开始进去房间.....
// 1 开始进去房间.....
// 2 已经进去房间.....
// 2 已经离开房间
// 1 已经进去房间.....
// 1 已经离开房间
// main 执行完毕