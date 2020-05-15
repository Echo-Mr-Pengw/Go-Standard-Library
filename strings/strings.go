// 用户操作字符串的包
package main

import (
	"fmt"
	"strings"
)

func main() {
	// EqualFold(s1, s2) 字符串的比较。比较s1和s2是否相等，不区分大小写
	fmt.Println("两个字符串的比较：", strings.EqualFold("Go", "go"))  // true

	// HasPrefix(s, prefix) 判断字符串s的前缀（第一个字符）是不是prefix。注意：区分大小写
	fmt.Println(strings.HasPrefix("abcd", "a"))    // true
	fmt.Println(strings.HasPrefix("Abcda", "a"))   // false

	// HasSuffix(s, suffix) 判断字符串的后缀（最后一个字符）是不是suffix。注意：区分大小写
	fmt.Println(strings.HasSuffix("abA", "a"))    // false
	fmt.Println(strings.HasSuffix("abA", "A"))    // true

	// Contains(s, sub) 字符串s是否包含子字符串sub。注意：区分大小写
	fmt.Println(strings.Contains("abcd", "B"))   // false
	fmt.Println(strings.Contains("abcd", "b"))   // true

	// ContainsAny(s, char) 字符串s是否包含char中的任意一个字符
	fmt.Println(strings.ContainsAny("abc", "a"))   // true
	fmt.Println(strings.ContainsAny("abc", "ac"))  // true
	fmt.Println(strings.ContainsAny("abc", "d"))   // false

	// Count(s, char) 返回char字符串在s字符串中出现的总个数
	fmt.Println(strings.Count("abcbbc", "bc"))  // 2
	fmt.Println(strings.Count("b", "aaa"))      // 0
	fmt.Println(strings.Count("aaaaa", ""))     // 6

	// Index(s, sub) 子字符串sub在字符串s中第一次出现的位置，如果不存在返回-1。
	// 注意：区分大小写
	fmt.Println(strings.Index("aaa", "a"))  //  0
	fmt.Println(strings.Index("baa", "a"))  //  1
	fmt.Println(strings.Index("baa", "A"))  //  -1

	// IndexByte(s, char) 字符char在字符串s中第一次出现的位置，不存在返回-1。
	// 注意：区分大小写；是字符而不是字符串。
	fmt.Println(strings.IndexByte("abc", 'a'))  // 0
	fmt.Println(strings.IndexByte("aca", 'a'))  // 0
	fmt.Println(strings.IndexByte("abc", 'A'))  // -1

	// IndexAny(s, char) 字符串char任意一个字符在s字符串中出现的位置，找到则返回，不会再继续匹配.
	// 整个字符串都扫描完，还没匹配到，返回-1。
	// 注意：区分大小写；从0开始
	fmt.Println(strings.IndexAny("abc", "dfc")) // 2
	fmt.Println(strings.IndexAny("abc", "C"))   // -1
	fmt.Println(strings.IndexAny("abc", "a"))   // 0

	// LastIndex(s, sub) 子串sub在字符串s中最后出现的位置，没有找到，返回-1
	// 注意：区分大小写；从0开始算，不是1。
	fmt.Println(strings.LastIndex("abcbc", "bc"))  // 3
	fmt.Println(strings.LastIndex("abcbc", "C"))   // -1
	fmt.Println(strings.LastIndex("abcbc", "c"))   // 4

	// LastIndexAny(s, char) char中任意一个字符在s中出现的位置，没有找到，返回-1
	// 注意：区分大小写；U以及第二种情况
	fmt.Println(strings.LastIndexAny("abc", "bd")) // 1
	fmt.Println(strings.LastIndexAny("abc", "bc")) // 2
	fmt.Println(strings.LastIndexAny("abc", "BC")) // -1
}

