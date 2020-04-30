package main
import (
	"fmt"
	"unsafe"
)

func main() {
	str := "ab"
	// Sizeof返回所占用的字节数
	fmt.Println(unsafe.Sizeof(str))
}
