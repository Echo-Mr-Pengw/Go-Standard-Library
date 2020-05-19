package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	str := "hello world"
	// 字符串强制转换成切片
	shaData := []byte(str)
	fmt.Printf("%x", sha256.Sum256(shaData))  // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
}