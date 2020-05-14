// 用户操作字符串的包
package main

import (
	"fmt"
	"strings"
)

func main() {
	// EqualFold(s1, s2) 字符串的比较。比较s1和s2是否相等，不区分大小写
	fmt.Println("两个字符串的比较：", strings.EqualFold("Go", "go"))  // true

	// HasPrefix(s, prefix) 判断字符串s的前缀（第一个字符）是不是prefix。注意：区分大小写
	fmt.Println(strings.HasPrefix("abcd", "a"))    // true
	fmt.Println(strings.HasPrefix("Abcda", "a"))   // false

	// HasSuffix(s, suffix) 判断字符串的后缀（最后一个字符）是不是suffix。注意：区分大小写
	fmt.Println(strings.HasSuffix("abA", "a"))    // false
	fmt.Println(strings.HasSuffix("abA", "A"))    // true
}

