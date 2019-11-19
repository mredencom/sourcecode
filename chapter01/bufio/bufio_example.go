package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

// 一些关于bufio的例子使用
func main() {
	//ReaderExample()
	//ReaderErrExample()
	//ReadBytesExample()
	//ReadStringExample()
	//ReadLineExample()
	PeekExample()
}

// ReadSlice的一个例子
func ReaderExample() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \n It is the home of gophers"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}

// ReadSlice的一个error例子
func ReaderErrExample() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)
	line, err := reader.ReadSlice('\n')
	fmt.Printf("line:%s\n terror:%s\n", line, err)
	line, err = reader.ReadSlice('\n')
	fmt.Printf("line:%s\n terror:%s\n", line, err)
}

// ReadBytes的一个例子🌰
func ReadBytesExample() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	n, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}

// ReadString 的一个例子🌰
func ReadStringExample() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadString('\n')
	fmt.Printf("the line:%s\n", line)
	n, _ := reader.ReadString('\n')
	fmt.Printf("the line:%s\n", n)
}

// Readline 的一个例子🌰
func ReadLineExample() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _, _ := reader.ReadLine()
	fmt.Printf("line:%s\n", line)
	n, _, _ := reader.ReadLine()
	fmt.Printf("linessss:%s\n", line)
	fmt.Println(string(n))
}

// peek是读取reader前n个字节
func PeekExample() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	go _peekE(reader)
	go reader.ReadBytes('\t')
	//time.Sleep(1 * time.Second)
	time.Sleep(1e8)
}
func _peekE(reader *bufio.Reader) {
	line, _ := reader.Peek(14)
	fmt.Printf("%s\n", line)
	//time.Sleep(1)
	fmt.Printf("%s\n", line)
}
