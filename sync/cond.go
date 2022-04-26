package main

import(
	"fmt"
	"sync"
	"time"
)

var done bool
var ch chan bool

func read(cond *sync.Cond) {
	for i := 0; i < 3; i++ {
		go func(n int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			fmt.Printf("goroutine%d start\n", n)
			if !done {
				cond.Wait()
			}
			fmt.Printf("goroutine%d end\n", n)
		}(i)
	}
}

//func test() {
//	go func() {
//		fmt.Println("test start")
//		n := <-ch
//		fmt.Println("test end ", n)
//	}()
//}

func write() {
	done = true
}

func main() {

	lock := sync.Mutex{}
	cond := sync.NewCond(&lock)

	fmt.Println("main start read")
	read(cond)
	time.Sleep(time.Second * 2)
	fmt.Println("main start write")

	// test()  // 对此goroutine无影响

	fmt.Println("main start broadcast")
	cond.Broadcast()  // 通知全部的堵塞的goroutine 退出

	// write()
	//cond.Signal()  // 谁先抢占锁，谁先退出
	time.Sleep(time.Second * 10)
}