package main

import "testing"

func TestTeardownParallel(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", parallelTest1)
		t.Run("Test2", parallelTest2)
	})
}

func parallelTest1(t *testing.T) {
	//t.Parallel() // 开启并行
	t.Log(111111)
	Fib(1)
}

func parallelTest2(t *testing.T) {
	//t.Parallel() // 开启并行
	t.Log(222222)
	Fib(3)
}
