package main

import (
	"bytes"
	"fmt"
)

//这是一个关于strings buffer的使用，这个是还是很是常用的  需要了解一下

func main() {
	NewBufferStringExample()
}

// 创建一个buffer对象
func NewBufferStringExample() {
	a := bytes.NewBufferString("Good Night")
	x, err := a.ReadBytes('h')
	if err != nil {
		fmt.Println("delim:t err:", err)
	} else {
		fmt.Println(string(x))
	}
	// 截取字符串 载入缓存中
	a.Truncate(1)
	fmt.Println(a.Len(), a.String())
	a.WriteString("Good Night")
	a.Truncate(5)
	fmt.Println(a.Len())
	y, err := a.ReadString('N')
	if err != nil {
		fmt.Println("delim:N err:", err)
	} else {
		fmt.Println(y)
	}

}
