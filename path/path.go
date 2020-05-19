// path包：与路径相关的包
package main

import (
	"fmt"
	"path"
)
func main() {

	// IsAbs(path)，path为绝对路径返回true，否则false
	fmt.Println(path.IsAbs("a/b/c"))   // false
	fmt.Println(path.IsAbs("/a/b/c"))  // true

	// Split(path) 通过path中的最后一个"/"进行判断， 返回目录和文件名
	fmt.Println(path.Split("a/b"))      // a/  b （目录为a/，文件为b）
	fmt.Println(path.Split("/a"))       // /   b （目录为/，文件为b）
	fmt.Println(path.Split("a"))        //     b （目录为空，文件为b）
	fmt.Println(path.Split("a/b/c"))    // a/b/ c （目录a/b/，文件为c）

	// Join(...s) 将对个字符串s，通/连接起来；空字符串会忽略。
	fmt.Println(path.Join("a", "b", "c"))     // a/b/c
	fmt.Println(path.Join("a", "b", "c", "", "d")) // a/b/c/d

	// Dir(path) 返回path的目录
	// 注意：
	// path = ""或者单个字符， 返回.
	fmt.Println(path.Dir("b/c"))    // b
	fmt.Println(path.Dir("/b/c"))   // /b
	fmt.Println(path.Dir(""))      // .
	fmt.Println(path.Dir("a"))     // .

	// Base(path) 返回路径的最后一个字符
	// 注意：
	// path = ""， 返回.
	fmt.Println(path.Base("a/b"))    // b
	fmt.Println(path.Base("a/b/"))   // b
	fmt.Println(path.Base(""))       // .
	fmt.Println(path.Base("c"))      // c

	// Ext(path) 返回path的后缀名。如果path最后一个斜杠后面有点，则返回点后面的字符串；
	// 反之，返回空
	fmt.Println(path.Ext("a/b/c.css")) // .css
	fmt.Println(path.Ext("a/b/c"))     //
	fmt.Println(path.Ext("a/b/c."))    // .
}
