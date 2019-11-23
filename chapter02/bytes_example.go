package main

import (
	"bytes"
	"fmt"
)

// 关于bytes-byte的方法操作以及使用
func main() {
	//RunesExample()
	//ReaderExample()
	BufferExample()
}

// byte转rune
func RunesExample() {
	//第一种装换会出现乱码
	b := []byte("你好，世界")
	for k, v := range b {
		fmt.Printf("%d:%s |", k, string(v))
	}
	// 建议使用这个一种转换 支持中文，和多种字符
	r := bytes.Runes(b)
	for k, v := range r {
		fmt.Printf("%d:%s|", k, string(v))
	}
}

// bytes->reader bytes关于实现的Reader接口的
func ReaderExample() {
	x := []byte("你好，世界")
	r1 := bytes.NewReader(x)
	d11 := make([]byte, len(x))
	nn, _ := r1.Read(d11)
	fmt.Println(nn, string(d11))

	// 读取数据至 b
	r2 := bytes.Reader{}
	r2.Reset(x)
	d22 := make([]byte, len(x))
	ns, _ := r2.Read(d22)
	fmt.Println(ns, string(d22))

	// UnreadRune 进度下标指向前一个字符，如果 r.i <= 0 返回错误，且只能在每次 ReadRune 方法后使用一次，否则返回错误。
	ch, size, _ := r1.ReadRune()
	fmt.Println(size, string(ch))
	_ = r1.UnreadRune()
	ch, size, _ = r1.ReadRune()
	fmt.Println(size, string(ch))
	_ = r1.UnreadRune()

	by, _ := r1.ReadByte()
	fmt.Println(by)
	_ = r1.UnreadByte()
	by, _ = r1.ReadByte()
	fmt.Println(by)
	_ = r1.UnreadByte()

	d1 := make([]byte, 6)
	n, _ := r1.Read(d1)
	fmt.Println(n, string(d1))

	d2 := make([]byte, 6)
	n, _ = r1.ReadAt(d2, 0)
	fmt.Println(n, string(d2))

	w1 := &bytes.Buffer{}
	_, _ = r1.Seek(0, 0)
	_, _ = r1.WriteTo(w1)
	fmt.Println(w1.String())

}

// Buffer类型
func BufferExample() {
	// new buffer string
	a := bytes.NewBufferString("Hello World")
	// new buffer byte
	b := bytes.NewBuffer([]byte("Hello World"))
	c := bytes.Buffer{}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
