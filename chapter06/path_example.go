package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	FilePathExample()
}

// file path的使用
func FilePathExample() {
	//  返回路径是否是一个绝对路径
	fmt.Println(filepath.IsAbs("/Users/jiangtao/go/src/sourcecode"))
	fmt.Println(filepath.IsAbs("./"))

	// 函数返回 path 代表的绝对路径，如果 path 不是绝对路径，会加入当前工作目录以使之成为绝对路径。
	// 因为硬链接的存在，不能保证返回的绝对路径是唯一指向该地址的绝对路径。在 os.Getwd 出错时，Abs
	// 会返回该错误，一般不会出错，如果路径名长度超过系统限制，则会报错。
	fmt.Println(filepath.Abs("/Users/jiangtao/go/src/sourcecode"))

	// Rel 函数返回一个相对路径，将 basepath 和该路径用路径分隔符连起来的新路径在词法上等价于 targpath。
	// 也就是说，Join(basepath, Rel(basepath, targpath)) 等价于 targpath。如果成功执行，返回值总是
	// 相对于 basepath 的，即使 basepath 和 targpath 没有共享的路径元素。如果两个参数一个是相对路径而
	// 另一个是绝对路径，或者 targpath 无法表示为相对于 basepath 的路径，将返回错误。
	fmt.Println(filepath.Rel("/home/polaris/studygolang", "/home/polaris/studygolang/src/logic/topic.go"))
	fmt.Println(filepath.Rel("/home/polaris/studygolang", "/data/studygolang"))

	// TODO 路径的拼接和切分
	// 对于一个常规`文件`路径，我们可以通过 Split 函数得到它的目录路径和文件名 返回结果是文件的绝对路径和`文件`的名字
	fmt.Println(filepath.Split("/Users/jiangtao/go/src/sourcecode/chapter06/path_example.go"))
	// 当当前文件没有目录时 就dir=""
	fmt.Println(filepath.Split("path_example.go"))
	// 相对路径到绝对路径的转变，需要经过路径的拼接。Join 用于将多个路径拼接起来，会根据情况添加路径分隔符。
	str := make([]string, 10)
	str = append(str, "Users")
	str = append(str, "go")
	str = append(str, "/")
	str = append(str, "path_example.go")
	fmt.Println(filepath.Join(str...))

	// Clean 函数通过单纯的词法操作返回和 path 代表同一地址的最短路径。
	// 它会不断的依次应用如下的规则，直到不能再进行任何处理：
	// 将连续的多个路径分隔符替换为单个路径分隔符
	// 剔除每一个 . 路径名元素（代表当前目录）
	// 剔除每一个路径内的 .. 路径名元素（代表父目录）和它前面的非 .. 路径名元素
	// 剔除开始于根路径的 .. 路径名元素，即将路径开始处的 /.. 替换为 /（假设路径分隔符是 /）
	// 返回的路径只有其代表一个根地址时才以路径分隔符结尾，如 Unix 的 / 或 Windows 的 C:\。
	// 如果处理的结果是空字符串，Clean 会返回 .，代表当前路径。
	fmt.Println(filepath.Clean("/Users/jiangtao/./go/src/sourcecode/chapter06/path_example.go"))

	// 在当前目录下创建一个 learn.txt 文件和一个 symlink 目录，在 chapter06 目录下对 learn.txt 建一个符号链接 learn.txt.1
	fmt.Println(filepath.EvalSymlinks("chapter06/learn.txt.1"))
	fmt.Println(os.Readlink("chapter06/learn.txt.1"))
}
