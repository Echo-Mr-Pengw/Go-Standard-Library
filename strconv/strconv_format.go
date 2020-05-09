package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 把bool类型的true或者false 转换成 字符串“true”或者“false”
	s := strconv.FormatBool(false)
	fmt.Printf("%T, %s\n", s, s) // string, false

	// 把整型类型的值(有符合) 转换成字符串
	s1 := strconv.FormatInt(+100, 10)
	fmt.Printf("类型：%T， 值：%s\n", s1, s1) //类型：string， 值：100

	// 把整型类型的值(无符号) 转换成字符串
	s2 := strconv.FormatUint(1, 10)
	fmt.Printf("类型：%T, 值：%s", s2, s2) // 类型：string, 值：1
}
