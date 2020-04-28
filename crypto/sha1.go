package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "sha1"
	sha1Data := []byte(str)
	fmt.Printf("%x", sha1.Sum(sha1Data))
}