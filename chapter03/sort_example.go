package main

import (
	"fmt"
	"sort"
)

//数据结构与算法
func main() {
	//SortExample()
	//SearchSortExample()
	//GuessingGame()
	//IntSliceExample()
	InterfaceSortExample()
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
func SearchSortExample() {
	x := 16
	// TODO:注意已经升序排序 这里的数组必须是升序排序好的数组
	s := []int{3, 6, 8, 11, 45}
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

// int类型的slice进行排序
func IntSliceExample() {
	s := []int{1, 23, 2, 4, 5, 4, 0}
	sort.Ints(s)
	fmt.Println("正序：", s)
	// TODO search方法必须是正序
	x := sort.SearchInts(s, 1)
	fmt.Println("正序找到数据：", x)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println("倒序：", s)
}

// 使用扩展的排序 interface
// TODO 这个是可以排序结构体的排序方式
func InterfaceSortExample() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"阿呆", 24},
		{"阿信", 75},
	}
	sort.Slice(people, func(i, j int) bool {
		//return people[i].Age < people[j].Age
		return people[i].Name < people[j].Name
	})

	fmt.Println("Sort by age or name:", people)
	fmt.Println(sort.SliceIsSorted(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	}))

	a := []int{2, 3, 4, 200, 100, 21, 234, 56}
	x := 21

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })           // 升序排序
	index := sort.Search(len(a), func(i int) bool { return a[i] >= x }) // 查找元素

	if index < len(a) && a[index] == x {
		fmt.Printf("found %d at index %d in %v\n", x, index, a)
	} else {
		fmt.Printf("%d not found in %v,index:%d\n", x, a, index)
	}
}
