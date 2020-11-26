package main

import (
	"fmt"
	"net/http"
)

func main() {

	// 规范化header域中的键名 例如，"accept-encoding"规范化为"Accept-Encoding"。
	headerKey := http.CanonicalHeaderKey("user-age")
	fmt.Println(headerKey) // User-Age

	// 解析http版本字符串
	mi, mo, b := http.ParseHTTPVersion("http1.1")
	fmt.Println(mi, mo, b)  // 0 0 false

	// 状态码对应的文本
	text := http.StatusText(200)
	fmt.Println(text)  // ok

	// header相关操作
	// 初始化header域，可以设置多个
	header := http.Header{"Date": {"2020-11-24"}, "age":{"12"}}

	// 	获取header头域中的date值
	fmt.Println(header.Get("Date"))  // 2020-11-24

	// 设置header
	header.Set("cookie", "demo")
	fmt.Println(header.Get("cookie"))  // demo

	// 获取一个不存在的header域的键，返回空
	fmt.Println(header.Get("username"))  // ""

	// header域中添加数据
	header.Add("sessionKey", "sessionVal")
	fmt.Println(header.Get("sessionKey"))  // sessionVal

	// 删除header域中的键值对
	header.Del("sessionKey")
	fmt.Println(header.Get("sessionKey"))  // ""

	// cookie相关操作
	// 初始化cookie
	cookie := http.Cookie{Name:"go", Value:"cookie"}
	// 序列化cookie
	fmt.Println(cookie.String())
}
