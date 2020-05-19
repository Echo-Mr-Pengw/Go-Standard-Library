// md5包实现了MD5哈希算法

package main
import (
	"crypto/md5"
	"fmt"
)

func main() {
	// 常量 md5字节数的大小
	fmt.Println(md5.BlockSize)    // 64

	// 常量 md5校验和的字节数
	fmt.Println(md5.Size)        // 16

	// 返回 md5的校验和
	str := "go"
	md5Data :=[]byte(str)
	fmt.Printf("%x", md5.Sum(md5Data))   // 34d1f91fb2e514b8576fab1a75a89a6b
}

