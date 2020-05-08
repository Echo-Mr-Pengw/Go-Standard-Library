// sync包中的互斥锁 Mutex
// 开启两个线程模拟窗口1和窗口2对10张票进行售卖
package main

import (
	"fmt"
	"sync"
)

// 定义全局变量
var tickets = 10
var w sync.WaitGroup
var m sync.Mutex

func main() {

	w.Add(2)
	go sellTicket("窗口1")
	go sellTicket("窗口2")
	w.Wait()
}

func sellTicket(name string) {
	defer w.Done()
	for {
		m.Lock()  // 给线程加锁
		if tickets > 0 {
			fmt.Println(name, "售出第", tickets, "张票")
			tickets--
		}else{
			 m.Unlock()  // 如果票售完，释放锁，以免造成死锁
			break
		}
		m.Unlock() // 线程执行完，释放锁
	}
	fmt.Println(name, "票全部售完")
}
