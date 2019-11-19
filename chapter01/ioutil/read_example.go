package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//ioutile中的read的扩展包

func main() {
	//dir := os.Args[1]

	fmt.Println(os.Args)
	ListAll("/Users/jiangtao/go/src/sourcecode", 0)
}

//ioutil 读取下所有的文件 读取指定目录下的文件
func ListAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|__\t")
			}
			fmt.Println(info.Name(), "\\")
			ListAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|__\t")
			}
			fmt.Println(info.Name())
		}
	}
}


