package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//CopyExample()
	//CopyNExample()
	CopyBufferExample()
}

// Copy 将 src 复制到 dst，直到在 src 上到达 EOF 或发生错误。它返回复制的字节数，如果有错误的话，还会返回在复制时遇到的第一个错误。
// 成功的 Copy 返回 err == nil，而非 err == EOF。由于 Copy 被定义为从 src 读取直到 EOF 为止，因此它不会将来自 Read 的 EOF 当做错误来报告。
// 若 dst 实现了 ReaderFrom 接口，其复制操作可通过调用 dst.ReadFrom(src) 实现。此外，若 src 实现了 WriterTo 接口，其复制操作可通过调用 src.WriteTo(dst) 实现。
// 将输出的打印到屏幕上
func CopyExample() {
	io.Copy(os.Stdout, os.Stdin)
	fmt.Println("Got EOF --bye")
}

// CopyN 将 n 个字节(或到一个error)从 src 复制到 dst。 它返回复制的字节数以及在复制时遇到的最早的错误。当且仅当err == nil时,written == n 。
// 若 dst 实现了 ReaderFrom 接口，复制操作也就会使用它来实现。
func CopyNExample() {
	io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 8)
}

func CopyBufferExample() {
	io.CopyBuffer(os.Stdout, os.Stdin, make([]byte,1))
}
