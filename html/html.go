// Go之html标准库
// 转义和解转义HTML文本

package main

import (
	"fmt"
	"html"
)

func main() {

	// EscapeString 转义HTML文本（<、>、&、'、"）
	str := "hello < wolrld > Golang&html'"
	escapeStr := html.EscapeString(str)
	fmt.Println("转义后的字符串：", escapeStr) //  hello &lt; wolrld &gt; Golang&amp;html&#39;

	// UnescapeString 解转义HTML文本
	unEscapeStr := html.UnescapeString(escapeStr)
	fmt.Println(unEscapeStr) // hello < wolrld > Golang&html'
}
