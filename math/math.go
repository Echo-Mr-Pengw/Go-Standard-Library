// math包提供了 常用的数学常量和数学函数
package  main

import (
	"fmt"
	"math"
)
func main() {

	// 常量PI
	fmt.Println(math.Pi)   // 3.141592653589793

	// Ceil(x) 返回大于x的最大数
	fmt.Println(math.Ceil(3.1))  // 4

	// Floor(x) 返回小于x的最小数
	fmt.Println(math.Floor(1.1))  // 1

	// Abs(x) 返回x的绝对值
	fmt.Println(math.Abs(-1))   // 1

	// Max(x, y) 返回x和y的最大值
	fmt.Println(math.Max(1.6, 2.0)) // 2

	// Min(x, y) 返回x和y的最小值
	fmt.Println(math.Min(2.1,  2.2))  // 2.1
}

