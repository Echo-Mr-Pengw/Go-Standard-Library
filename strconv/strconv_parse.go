package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 字符串参数转换成bool值
	// 参数：1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	b, err := strconv.ParseBool("0")
	fmt.Println(b, err) // false nil

	// 字符串转换成整型，第一个参数可以是正负数
	i, err := strconv.ParseInt("-11", 10, 0)
	fmt.Println(i, err)  // -11 nil

	// 字符串转换成无符号的整型，第一个参数是无符号的
	i1, err := strconv.ParseUint("12", 10, 0)
	fmt.Println(i1, err)  // 12 nil

	// 字符串的浮点型转换成浮点型
	f, err := strconv.ParseFloat("1.20", 64)
	fmt.Println(f, err) // 1.2 nil


}

