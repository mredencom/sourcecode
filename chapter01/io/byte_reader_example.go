package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// 字节读取
func main() {
	//ByteReaderByte()
	//RuneReaderReader()
	ReaderLimited()
}

func ByteReaderByte() {
	var ch byte
	fmt.Println("输入：")
	fmt.Scanf("%c\n", &ch)
	buffer := new(bytes.Buffer)
	err := buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("写入一个字节成功，我们准备读取---")
		newCh, _ := buffer.ReadByte()
		fmt.Printf("读取数据:%c\n", newCh)
	} else {
		fmt.Println(err)
	}
}

// strings.Index
func RuneReaderReader() {
	n := strings.Index("行业交流群", "交流")
	//nu:=strings.ToValidUTF8()
	fmt.Println(n)
}

func ReaderLimited() {
	content := "This Is LimitReader Example"
	reader := strings.NewReader(content)
	limitReader := &io.LimitedReader{R: reader, N: 8}
	for limitReader.N > 0 {
		tmp := make([]byte, 2)
		fmt.Print("tmp:", tmp,"\n")
		fmt.Println()
		limitReader.Read(tmp)
		fmt.Printf("%s", tmp)
	}
}
