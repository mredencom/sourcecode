package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)

// 管道的读取和写入 pipeWriter 写入管道书阻塞的 必须pipeReader读取完才可以继续写入
// pipeWriter和pipeReader属于同步 不能在一个goroutine中执行
func main() {
	pipeReader, pipeWriter := io.Pipe()
	go WriterPipe(pipeWriter)
	go ReaderPipe(pipeReader)
	time.Sleep(30 * time.Second)
}

// 管道写入
func WriterPipe(writer *io.PipeWriter) {
	data := []byte("Go语言中文网制作")
	for i := 0; i <= 3; i++ {
		n, err := writer.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("写入字节：%d\n", n)
	}
	_ = writer.CloseWithError(errors.New("写入关闭"))
}

// 管道读取
func ReaderPipe(reader *io.PipeReader) {
	buf := make([]byte, 128)
	for {
		fmt.Println("接口端开始阻塞5秒...")
		time.Sleep(5 * time.Second)
		fmt.Println("现在开始接受数据")
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("收到字节: %d\n buf内容: %s\n", n, buf)
	}
}
