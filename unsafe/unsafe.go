package main
import (
	"fmt"
	"unsafe"
)

func main() {
	str := "ab"
	// 类型本身数据所占用的字节数
	fmt.Println(unsafe.Sizeof(str))

	// 返回类型的对齐方式（即类型v在内存中占用的字节数）
	fmt.Println(unsafe.Alignof(str))

	// 它返回该结构起始处与该字段起始处之间的字节数。
	// fmt.Println(unsafe.Offsetof(str))
}
