package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	str := "hello world"
	// 字符串强制转换成切片
	shaData := []byte(str)
	fmt.Printf("%x", sha256.Sum256(shaData))
}