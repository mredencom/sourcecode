package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	TeeReaderExample()
}
// TeeReader 返回一个 Reader，它将从 r 中读到的数据写入 w 中。所有经由它处理的从 r 的读取都匹配于对应的对 w 的写入。
// 它没有内部缓存，即写入必须在读取完成前完成。任何在写入时遇到的错误都将作为读取错误返回。
func TeeReaderExample() {
	reader := io.TeeReader(strings.NewReader("我是tee命令"), os.Stdout)
	reader.Read(make([]byte, 20))
}
