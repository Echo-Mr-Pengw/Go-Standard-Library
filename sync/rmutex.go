// RWMutex 读写锁
// RLock   读锁。有写操作需要等到读锁释放；有读操作不需要等待。
// RUnlock 释放读加锁。在未释放读锁之前，所有的写操作需要等待。
// Lock    写锁。加写锁后，不管是读还是写操作，必须等待写锁被释放。
// Unlock  释放写锁

// 此案例介绍读锁

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var rmutex = sync.RWMutex{}

func main() {

	wg.Add(2)

	go readData(1)
	go readData(2)

	wg.Wait()

	fmt.Println("main 结束...")
}

func readData(name int) {
	defer wg.Done()

	fmt.Println(name, "准备开始读数据......")
	// 加读锁，不影响其他协程的读操作。如果有读操作，则需等待
	rmutex.RLock()
	fmt.Println(name, "正在读数据")
	// 休眠
	time.Sleep(time.Second)
	fmt.Println(name, "数据读完")
	// 释放读锁
	rmutex.RUnlock()
}

// 通过下面的打印结果可以看出，
// 1虽然加读锁，其正在读数据，还未释放锁（1 数据还没读完）；
// 此时，2也可以读数据

// output:
// 1 准备开始读数据......
// 1 正在读数据
// 2 准备开始读数据......
// 2 正在读数据
// 2 数据读完
// 1 数据读完
// main 结束...