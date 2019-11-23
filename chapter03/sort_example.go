package main

import (
	"fmt"
	"sort"
)

//数据结构与算法
func main() {
	//SortExample()
	//SearchSortExample()
	GuessingGame()
}

//排序方式的实现
func SortExample() {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
		{"leao1", 90},
	}
	// 打印未排序的 stus 数据
	fmt.Println("默认顺序:\n\t", stus)
	//StuScores 已经实现了 sort.Interface 接口 , 所以可以调用 Sort 函数进行排序
	sort.Sort(stus)
	// 判断是否已经排好顺序，将会打印 true
	fmt.Println("是否排序?\n\t", sort.IsSorted(stus))
	// 打印排序后的 stus 数据
	fmt.Println("升序后:\n\t", stus)

	sort.Sort(sort.Reverse(stus))
	fmt.Println("降序后:\n\t", stus)
}

// 学生的一个结构体
type StuScore struct {
	name  string
	score int
}

type StuScores []StuScore

// 实现排序第一个方法 Len
func (s StuScores) Len() int {
	return len(s)
}

// 实现排序的第二个方法 Less
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

// 实现排序第三个方法 swap交换方法
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 这里数组必须是排序的
func SearchSortExample()  {
	x := 16
	s := []int{3, 6, 8, 11, 45} // 注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, " 在 s 中的位置为：", pos)
	} else {
		fmt.Println("s 不包含元素 ", x)
	}
}
// 查询
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}