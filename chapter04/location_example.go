package main

import (
	"fmt"
	"time"
)

func main() {
	LocationExample()
}

//时区
func LocationExample() {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic("出错了")
	}
	fmt.Println(location)
}
