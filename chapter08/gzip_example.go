package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	GzipWriterExample()
}

func GzipWriterExample() {
	file, errs := os.Create("chapter08/jt.gzip")
	defer file.Close()
	if errs != nil {
		log.Fatal(errs)
	}
	//这边还是写入缓存buf中数据
	//var buf bytes.Buffer
	zw := gzip.NewWriter(file)
	defer zw.Close()
	// 这个时候我们需要设置压缩头 Header 这个相当于一个http协议header
	fr, err := os.Open("chapter08/jt.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	//获取这个文件的信息 提供下面使用
	fi, err := fr.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	zw.Name = fr.Name()
	zw.ModTime = time.Date(2019, time.September, 11, 11, 11, 11, 0, time.Local)
	zw.Comment = "This is created file on system"
	// 打印基本的属性信息
	fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zw.Name, zw.Comment, zw.ModTime.Local())
	//建立一个缓冲区
	buf := make([]byte, fi.Size())
	// 把数据读取到缓冲区，
	_, err = fr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	// 取出缓冲区的数据 写入到gzip
	_, err = zw.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
}
