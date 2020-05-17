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

	// Title(s) 以空格为基准，将s字符串（英文字符串）,的首字母转换成大写。
	fmt.Println(strings.Title("ab cd")) // Ab Cd
	fmt.Println(strings.Title("abcd"))  // Abcd
	fmt.Println(strings.Title("abcd Aa"))  // Abcd Aa

	// ToTitle(s) 将字符串s中的每个英文字母转成其对应的大写字母
	fmt.Println(strings.ToTitle("Ab Cd"))  // AB CD
	fmt.Println(strings.ToTitle("abCd"))   // ABCD

	// ToLower(s) 将字符串s中每个英文字符都转换成其对应的小写字母
	fmt.Println(strings.ToLower("ABC DE")) // abc de
	fmt.Println(strings.ToLower("AbcDe"))  // abcde
	fmt.Println(strings.ToLower("A中Bc"))  // a中bc

	// ToUpper(s) 将字符串s中的每个英文字母都转换成其对应的大写字母
	fmt.Println(strings.ToUpper("abc de"))   // ABC DE
	fmt.Println(strings.ToUpper("AbcDe"))    // ABCDE
	fmt.Println(strings.ToUpper("Go语言Hi")) // GO语言HI

	// Repeat(s, count) 返回count个s。
	// 注意：如果count必须大于等于0；
	// count = 0   返回空字符串
	// count = 1   返回字符串s本身
	// count = -1  count 小于0 报错panic
	fmt.Println(strings.Repeat("Ab", 2))      // AbAb
	fmt.Println(strings.Repeat("ab cd", 2))   // ab cdab cd
	fmt.Println(strings.Repeat("a", 0))       // ""
	// fmt.Println(strings.Repeat("v", -1))               // 发生panic

	// Replace(s, old, new, n) 通过new字符串参数取代s字符串中前n个与old字符串相同的字符。
	// 注意：n的取值
	// n = 0：返回原s字符串
	// n < 0：取代s字符串中与old字符相同的全部字符，与n参数设置的数量无关。
	fmt.Println(strings.Replace("ab abc abcd", "ab", "新", 2))  // 新 新c abcd
	fmt.Println(strings.Replace("ab abc abcd", "ab", "新", 4))  // 新 新c 新cd
	fmt.Println(strings.Replace("ab abc abcd", "ab", "新", 0))  // ab abc abcd
	fmt.Println(strings.Replace("ab abc abcd ab", "ab", "新", -1))  // 新 新c 新cd 新
	fmt.Println(strings.Replace("ab abc", "e", "新", 2))       // ab abc

	// Trim(s, cutset) 去除字符串s前后的cutset
	// 注意：
	// 如果cutset在s字符串前后不存在，则返回原s。
	// 如果cutset="", 则返回原s。
	// 如果cutset=" ", 则去除s字符串前后的空格。
	fmt.Println(strings.Trim(" abc ", " "))       // abc
	str := " ab cd "
	fmt.Println(len(str), len(strings.Trim(str, ""))) // 7  7
	fmt.Println(strings.Trim("abc", "e"))         // abc

	// TrimLeft(s, cutset) 去除字符串s左边的cutset
	fmt.Println(strings.TrimLeft("abc", "a"))   // bc

	// TrimRight(s, cutset) 去除字符串s右边的cutset
	fmt.Println(strings.TrimRight("abcd", "d")) // abc

	// Fields(s) 以空格切割字符串s，返回切片字符串
	// 注意：
	// 字符串s前后有无空格，返回的值都是一样的
	fmt.Println(strings.Fields("a b c d"))   // [a b c d]
	fmt.Println(strings.Fields(" a b c d ")) // [a b c d]
	fmt.Println(strings.Fields(" "))         // []

	// Join(s, sep) 通过sep将切片字符串s连接成字符串
	s := []string{"Go", "PHP"}
	fmt.Println(strings.Join(s, "-")) // Go-PHP

	// 用sep将字符串s切割，返回切片字符串
	// 注意：
	// 如果sep在字符串s中不存在 返回原s
	//如果以字符串s中的首尾字母作为sep切割是，注意
	fmt.Println(strings.Split("a-b-c", "-")) // [a b c]
	fmt.Println(strings.Split("a-b-c", "c")) // [a b ]
	fmt.Println(strings.Split("a-b-c", "a")) // [ b c]
}

