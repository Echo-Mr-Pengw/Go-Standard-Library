package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "sha1"
	sha1Data := []byte(str)
	fmt.Printf("%x", sha1.Sum(sha1Data))   // 415ab40ae9b7cc4e66d6769cb2c08106e8293b48
}