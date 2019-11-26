package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//FileExample()
	//FileExample02()
	ReadAndOutputDir("./", 3)
}

// O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
// O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
// O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
// O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
// O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
// O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE 配合使用，文件必须不存在
// O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步 I/O
// O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
// file 的操作
// TODO 关闭一个文件时需要注意：1. 关闭一个未打开的文件；2. 两次关闭同一个文件；

func FileExample() {
	file, err := os.Open("chapter06/doc_example.go")
	if err != nil {
		log.Fatal(err)
	}
	b := make([]byte, 100)
	file.Read(b)
	fmt.Println(string(b))
	defer file.Close()
}

func FileExample02() {
	file, err := os.Create("chapter06/learn.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileMode := getFileMode(file)
	log.Println("file mode:", fileMode)
	//
	file.Chmod(fileMode | os.ModeSticky)

	log.Println("change after, file mode:", getFileMode(file))
}

func getFileMode(file *os.File) os.FileMode {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.Mode()
}

//读取目录下所有的
func ReadAndOutputDir(rootPath string, deep int) {
	file, err := os.Open(rootPath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()

	for {
		fileInfos, err := file.Readdir(100)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("readdir error:", err)
			return
		}

		if len(fileInfos) == 0 {
			break
		}

		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				if deep > 0 {
					ReadAndOutputDir(filepath.Join(rootPath, string(os.PathSeparator), fileInfo.Name()), deep-1)
				}
			} else {
				fmt.Println("file:", fileInfo.Name(), "in directory:", rootPath)
			}
		}
	}
}
