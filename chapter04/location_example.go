package main

import "time"

func main() {

}

//时区
func LocationExample() {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic("出错了")
	}
	location

}
