package main

import (
	"fmt"
	"time"
)

func main() {
	//TimeExample()
	TimeExample02()
}

func TimeExample() {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-11-25 09:14:00", time.Local)
	// 计算两个时间小时差
	fmt.Println(time.Now().Sub(t).Hours())
	fmt.Println(time.Now().Format("2006-01-02 03:04:05"))
}

func TimeExample02() {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
	// 整点（小时向下取整）类似小数取整
	fmt.Println(t.Truncate(1 * time.Hour))
	// 整点（小时最接近）类似小数取整
	fmt.Println(t.Round(1 * time.Hour))

	// 整点（分钟向下取整）类似小数取整
	fmt.Println(t.Truncate(1 * time.Minute))
	// 整点（分钟最接近）类似小数取整
	fmt.Println(t.Round(1 * time.Minute))

	// 整点（秒向下取整）类似小数取整
	fmt.Println(t.Truncate(1 * time.Second))
	// 整点（秒最接近）类似小数取整
	fmt.Println(t.Round(1 * time.Second))

	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
	fmt.Println(t2)

}
