package main

import (
	"fmt"
	"strings"
)

// 字符串的操作
func main() {
	//CompareExample("a", "c")
	//ContainsExample("hello world1", "1")
	CountExample()
}

// 字符串比较
// a<b 返回-1
// a>b 返回1
// a=b 返回1
func CompareExample(a, b string) {
	// Compare 函数，用于比较两个字符串的大小，如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -1 ，反之返回 1 。
	// 不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观。
	n := strings.Compare(a, b)
	fmt.Println(n)
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))

	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println(strings.EqualFold("壹", "一"))
}

// 是否存在某个字符或子串
func ContainsExample(s, substr string) {
	contains := strings.Contains(s, substr)
	fmt.Println("contains：", contains)
	containsAny := strings.ContainsAny(s, substr)
	fmt.Println("containsAny：", containsAny)
	containsRune := strings.ContainsRune(s, 1)
	fmt.Println("containsRune：", containsRune)
	lastIndex := strings.LastIndex(s, substr)
	fmt.Println("lastIndex：", lastIndex)
	indexBytes := strings.IndexByte(s, byte(1))
	fmt.Println("indexBytes：", indexBytes)
}

// 检查字符串出现的次数
func CountExample() {
	s := "hello world"
	sep := "ll"
	n := strings.Count(s, sep)
	fmt.Println(n)
}
