package main

import(
	"fmt"
	"sync"
)

var once sync.Once
var synWg sync.WaitGroup

func main() {

	synWg.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Printf("goroutine%d, start\n", n)
			once.Do(func() {
				fmt.Printf("goroutine%d, once.Do\n", n)  // 只会被执行一次
			})
			fmt.Printf("goroutine%d, end\n", n)
			synWg.Done()
		}(i)
	}
	synWg.Wait()
}