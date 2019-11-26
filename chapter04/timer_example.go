package main

import (
	"fmt"
	"time"
)

// 定时器的使用
func main() {
	//TimerExample01()
	TimerExample02()
}

// timer定时器的使用
func TimerExample01() {
	c := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		//time.Sleep(3 * time.Second)
		<-c
	}()
	select {
	case c <- 1:
		fmt.Println("channel...")
	case <-time.After(2 * time.Second):
		close(c)
		fmt.Println("timeout...")
	}
}

// timer 使用2
func TimerExample02() {
	start := time.Now()
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("这里产生函数回调, elaspe:", time.Now().Sub(start))
	})
	//time.Sleep(1 * time.Second)
	// Reset 在 Timer 还未触发时返回 true；触发了或 Stop 了，返回 false
	if timer.Reset(3 * time.Second) {
		fmt.Println("在这里触发了timer中的stopTimer")
	} else {
		fmt.Println("在这里没有触发了timer中的stopTimer")
	}
	time.Sleep(10 * time.Second)
}
