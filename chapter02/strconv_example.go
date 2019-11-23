package main

import (
	"fmt"
	"strconv"
	"time"
)

// 字符串的基本类型装换
// strconv — 字符串和基本数据类型之间转换
// 这里的基本数据类型包括：布尔、整型（包括有 / 无符号、二进制、八进制、十进制和十六进制）和浮点型等。
func main() {
	//s := IntToStringExample(2)
	//fmt.Println(s)
	//IntToStringExample02()
	//ParseBoolExample()
	//ParseFloatExample()
	QuotaExample()
}

func IntToStringExample(base int) string {
	const digits = "0123456789abcdefghijklmnopqrstuvwxyz"
	u := uint64(127)
	var a [65]byte
	i := len(a)
	b := uint64(base)
	for u >= b {
		i--
		a[i] = digits[uintptr(u%b)]
		u /= b
	}
	i--
	a[i] = digits[uintptr(u)]
	return string(a[1:])
}

func IntToStringExample02() {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		fmt.Sprintf("%d", i)
	}
	fmt.Println(time.Now().Sub(startTime))

	startTime = time.Now()
	for i := 0; i < 10000; i++ {
		strconv.Itoa(i)
	}
	fmt.Println(time.Now().Sub(startTime))
}

//  字符串和布尔值之间的转换
func ParseBoolExample() {
	// 接受 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False 等字符串；
	// 其他形式的字符串会返回错误
	bt, _ := strconv.ParseBool("0")
	fmt.Println(bt)
	// 直接返回 "true" 或 "false"
	bstr := strconv.FormatBool(true)
	fmt.Println(bstr)
	// 将 "true" 或 "false" append 到 dst 中
	// 这里用了一个 append 函数对于字符串的特殊形式：append(dst, "true"...)
	dst := make([]byte, 0)
	sb := strconv.AppendBool(dst, true)
	fmt.Println(string(sb))
}

// 字符串和浮点数之间的转换
func ParseFloatExample() {
	f, _ := strconv.ParseFloat("1212", 32)
	fmt.Println(f)
	str1 := strconv.FormatFloat(1223.13252, 'e', 3, 32) // 结果：1.223e+03
	str2 := strconv.FormatFloat(1223.13252, 'g', 3, 32) // 结果：1.22e+03
	fmt.Println(str1, str2)
}

// 其它的导出函数 引用函数
func QuotaExample() {
	fmt.Println("This is", strconv.Quote("studygolang.com"), "website")
}
