package main

import (
	"fmt"
	"math"
)

func main() {
	SinExample()
}

// 正弦函数，反正弦函数，双曲正弦，反双曲正弦
func SinExample() {
	fmt.Println(math.Sin(30))
	fmt.Println(math.Asin(112))
	fmt.Println(math.NaN())
	math.Float32bits()
}
