package main

import (
	"fmt"
	"bytes"
)
func main() {
	a := []byte{'a'}
	b := []byte{'a', 'b'}

	// 比较字符的大小 a=b 返回0  a<b 返回-1  a>b 返回·
	fmt.Println(bytes.Compare(a, b)) // -1

	// 比较两个切片的内容是否相等
	fmt.Println(bytes.Equal(a, b))  // false

	fmt.Println(bytes.EqualFold(a, b))

	// 判断是否含有前缀切片(b的前缀切片是a)
	fmt.Println(bytes.HasPrefix(b, a))   // true

	// 判断是否含有后缀切片(b的后缀切片是a)
	fmt.Println(bytes.HasSuffix(b, a))   // false

	// 判断b切片中是否含有子切片a
	fmt.Println(bytes.Contains(b, a))  // true

	// 判断b中含有多少个a切片子串
	fmt.Println(bytes.Count(b, a))  // 1

	// a子切片第一次出现在b中的位置；不存在返回-1
	fmt.Println(bytes.Index(b, a))  // 0

	// 字符在b切片中第一次出现的位置，不存在返回-1
	fmt.Println(bytes.IndexByte(b, 'b'))  // 1

	// 将字符转换成大写的拷贝，之前的变量值不变
	fmt.Println(string(bytes.ToUpper(b)), string(b))  // AB ab
}
