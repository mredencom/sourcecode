package main

import "testing"

// 简单的快速测试使用
//func TestFib(t *testing.T) {
//	var (
//		in       = 7
//		expected = 13
//	)
//	actual := Fib(in)
//	if actual != expected {
//		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
//	}
//}

// 简单快速的测试2
func TestFib(t *testing.T) {
	var TestFib = []struct {
		in       int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
	}
	for _, tt := range TestFib {
		actual := Fib(tt.in)
		if actual != tt.expected {
			// 运行到t.Errorf 会终止运行
			t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}

// 做一个压力测试
func benchmarkFib(i int, b *testing.B) {
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
