package main

import (
	"bufio"
	"fmt"
	"os"
)

// Scanner 的基本使用
func main() {
	ScannerExample()
}

func ScannerExample() {
	//file, err := os.Create("chapter01/bufio/01.txt")
	file, err := os.Open("chapter01/bufio/01.txt")
	if err != nil {
		// 这里好像发生错误了
		panic(err)
	}
	// 文件用完之后需要关闭
	defer file.Close()
	file.WriteString("http://studygolang.com.\nIt is the home of gophers.\nIf you are studying golang, welcome you!")
	// 我们将file的offset设置文件的开头
	file.Seek(0, os.SEEK_SET)
	//然后我扫描读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
