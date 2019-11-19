package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

//MultiReader and MultiWriter use
func main() {
	//MultiReaderExample()
	MultiWriterExample()
}

// 代码中首先构造了一个 io.Reader 的 slice，由 strings.Reader 和 bytes.Buffer 两个实例组成，
// 然后通过 MultiReader 得到新的 Reader，循环读取新 Reader 中的内容。从输出结果可以看到，
// 第一次调用 Reader 的 Read 方法获取到的是 slice 中第一个元素的内容……也就是说，MultiReader
// 只是逻辑上将多个 Reader 组合起来，并不能通过调用一次 Read 方法获取所有 Reader 的内容。在所有的
// Reader 内容都被读完后，Reader 会返回 EOF。
func MultiReaderExample() {
	readers := []io.Reader{
		strings.NewReader("form strings reader 你好 "),
		bytes.NewBufferString("from buffer string"),
	}
	reader := io.MultiReader(readers...)
	data := make([]byte, 0, 128)
	buf := make([]byte, 10)

	for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
		if err != nil {
			panic(err) //出现错误了
		}
		data = append(data, buf[:n]...)
	}
	fmt.Printf("%s\n", data)
}

// 这段程序执行后在生成 tmp.txt 文件，同时在文件和屏幕中都输出：Go语言中文网。这和 Unix 中的 tee 命令类似。
func MultiWriterExample() {
	file, err := os.Create("chapter01/tmp.txt")
	if err != nil {
		// 出错了
		panic(err)
	}
	defer file.Close()
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	writer.Write([]byte("我是写入文件了内容了"))
}
