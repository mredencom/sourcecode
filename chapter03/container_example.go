package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

// 常见的数据结构使用
func main() {
	//HeapExample()
	//ListExample()
	RingExample()
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 堆的使用
func HeapExample() {
	h := &IntHeap{2, 1, 5}
	// Init 做了排序 down 就是把大的元素下沉
	heap.Init(h)
	fmt.Println("Init:", h)
	// 放一个元素进去
	heap.Push(h, 3)
	fmt.Println("Push:", h)
	// 删除最小的元素
	heap.Pop(h)
	fmt.Println("Pop:", h)
}

// 源代码提供的LIST(链表)使用
// TODO 链表的使用
func ListExample() {
	// 生成一个链表
	lists := list.New()
	lists.PushBack(1)
	e2 := lists.PushBack(2)
	lists.PushBack(3)
	fmt.Printf("len: %v\n", lists.Len())
	fmt.Printf("first: %#v\n", lists.Front())
	fmt.Printf("second: %#v\n", lists.Front().Next())
	fmt.Printf("second: %#v\n", lists.Back())
	fmt.Printf("second: %#v\n", lists.InsertAfter(30, e2))
}

// TODO 环链表的使用
func RingExample() {
	// 初始化一个环链表 建立一个长度为3环链表
	ring := ring.New(3)
	for i := 1; i <= 3; i++ {
		ring.Value = i
		ring = ring.Next()
	}
	//ring.Unlink(2)
	s := 0
	ring.Do(func(p interface{}) {
		s += p.(int)
	})
	fmt.Println("sum is ", s)
	fmt.Println("环的长度 ", ring.Len())
	fmt.Println("当前元素的下一个元素 ", ring.Next().Value)
	fmt.Println("当前元素的上一个元素 ", ring.Prev().Value)
	fmt.Println("当前元素 ", ring.Value)
}
