package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
)

// zip 例子合集
func main() {
	//ZipWriterExample()
	//ZipReaderExample()
	ExampleWriter_RegisterCompressor()
}

// 压缩
func ZipWriterExample() {
	// 创建一个待写入的缓存区
	buf, _ := os.Create("chapter08/zipJt.zip")
	// 创建一个文档
	w := zip.NewWriter(buf)
	// 增加一些例子数据
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "这个是一个示例数据介绍"},
		{"gopher.txt", "I'm gopher"},
		{"todo.txt", "我还有一些任务写入任务没有做完"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	// 关闭写入流 并且检查是否出错
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
}

// zip 读取压缩
func ZipReaderExample() {
	r, err := zip.OpenReader("chapter08/zipJt.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()
	// 我们这里准备做下迭代下数据
	for _, f := range r.File {
		// 打印下读取的文件名字
		fmt.Println(f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		// 把内容copy到控制台输出
		_, err = io.CopyN(os.Stdout, rc, 688)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}

func ExampleWriter_RegisterCompressor() {
	// Override the default Deflate compressor with a higher compression level.

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Register a custom Deflate compressor.
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// Proceed to add files to w.
}