package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	//ReadAtAt()
	//WriterAtAt()
	//ReadFormForm()
	//WriteToTo()
	SeekSeek()
}

// ReadAt接口方法的使用
func ReadAtAt() {
	reader := strings.NewReader("go语言中文网")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 20)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}

// WriterAt 接口方法
func WriterAtAt() {
	file, err := os.Create("chapter01/writeAt.txt")
	if err != nil {
		//创建文件失败
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余")
	n, err := file.WriteAt([]byte("go语言中文网"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

// ReadForm 接口方法的实现
func ReadFormForm() {
	// 打开writeAt.txt文件
	file, err := os.Open("chapter01/writeAt.txt")
	if file == nil {
		panic("文件不存在")
	}
	defer file.Close()
	//打开一个文件可能出现一个错误信息
	if err != nil {
		panic(err)
	}
	// 使用os.Stdout是为了输出到控制台
	writer := bufio.NewWriter(os.Stdout)
	n, err := writer.ReadFrom(file)
	//将所有的缓冲写入io.Writer
	_ = writer.Flush()
	fmt.Println("\n", n)

}

// WriteTo接口使用
func WriteToTo() {
	buf := bytes.NewBuffer([]byte("Go语言中文网"))
	n, err := buf.WriteTo(os.Stdout)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n", n)
}

// Seek取出指定开始位置偏移量
func SeekSeek() {
	reader := strings.NewReader("GO语言中文网")
	reader.Seek(-6, os.SEEK_END)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}
