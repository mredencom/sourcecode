package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ReaderExample()
}
//Reader接口
// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.1.html

func ReaderExample() {
FOREND:
	for {
		readerMenu()
		var ch string
		fmt.Scanln(&ch)
		var (
			data []byte
			err  error
		)
		switch strings.ToLower(ch) {
		case "1":
			fmt.Println("请输入不多于9个字符，以回车结束:")
			data, err = ReadForm(os.Stdin, 11)
		case "2":
			file, err := os.Open("/Users/jiangtao/go/src/sourcecode/chapter01/01.txt")
			if err != nil {
				fmt.Println("打开文件 01.txt 错误", err)
				continue
			}
			data, err = ReadForm(file, 90)
			file.Close()
		case "3":
			data, err = ReadForm(strings.NewReader("from string"), 12)
		case "4":
			fmt.Println("没有实现")
		case "b":
			fmt.Println("返回上级菜单")
			break FOREND
		case "q":
			fmt.Println("程序退出")
			os.Exit(0)
		default:
			fmt.Println("输入错误")
			continue
		}
		if err != nil {
			fmt.Println("数据读取失败，可以试试从其他输入源读取！")
		} else {
			fmt.Printf("读取到的数据是：%s\n", data)
		}
	}
}

// 根据 Go 语言中关于接口和满足了接口的类型的定义（Interface_types），
// 我们知道 Reader 接口的方法集（Method_sets）只包含一个 Read 方法，
// 因此，所有实现了 Read 方法的类型都满足 io.Reader 接口，也就是说，
// 在所有需要 io.Reader 的地方，可以传递实现了 Read() 方法的类型的实例。
func ReadForm(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], err
	}
	return p, err
}

//读取打印的信息
func readerMenu() {
	fmt.Println("")
	fmt.Println("*******从不同来源读取数据*********")
	fmt.Println("*******请选择数据源，请输入：*********")
	fmt.Println("1 表示 标准输入")
	fmt.Println("2 表示 普通文件")
	fmt.Println("3 表示 从字符串")
	fmt.Println("4 表示 从网络")
	fmt.Println("b 返回上级菜单")
	fmt.Println("q 退出")
	fmt.Println("***********************************")
}
