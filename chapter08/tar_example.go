package main

import (
	"archive/tar"
	"log"
	"os"
	"time"
)

// tar 压缩
func main() {
	TarExample()
}

func TarExample() {
	//var buf bytes.Buffer
	file, _ := os.Create("chapter08/jtTar.tar")
	tw := tar.NewWriter(file)
	// 这个建立三个文件 打包压缩
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "我想介绍一下我自己，你看怎么样"},
		{"gopher.txt", "Gopher names:\nJames"},
		{"todo.txt", "这个还是有的问题的"},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name:    file.Name,
			Mode:    0600,
			Size:    int64(len(file.Body)),
			ModTime: time.Date(2019, time.November, 11, 11, 11, 11, 11, time.Local),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	// 关闭一个压缩包的流
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}
}
