// builtin标准库包含Go的内建函数

package main

import "fmt"

func main() {

	// new(Type) Type必须是指针、通道、函数、接口、映射或切片；返回指针
	var p *int
	p = new(int)
	fmt.Println(*p)

	// make 分配并初始化一个类型为切片、映射、或通道的对象
	c := make([]int, 2)
	fmt.Println(c) // output：[0 0]

	// len(V) 返回V的长度
	fmt.Println(len(c)) // output：2

	// cap(V) 返回V的容量
	fmt.Println(cap(c)) // output：2

	// append 将1个或者多个元素追加到切片的末尾， 返回一个切片
	c1 := append(c, 1) // 一次性追加多个元素append(c, 1,2,3,4)
	fmt.Println(c1)    // [0 0 1]

	// copy 从原切片复制到目标切片，返回复制元素的个数
	c2 := make([]int, 2)
	copyNum := copy(c2, c1)
	fmt.Println(c2, copyNum) // [0 0] 2

	// delete按照指定的键将元素从映射中删除。若m为nil或无此元素，delete不进行操作。
	m := map[string]string{
		"name": "new1024kb",
		"age":  "20",
	}
	fmt.Println(m) // output: map[age:20 name:new1024kb]
	delete(m, "name")
	fmt.Println(m) // output: map[age:20]
}
