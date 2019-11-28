package main

import (
	"bytes"
	"html/template"
	"testing"
)

func BenchmarkFib100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

var bench = []struct {
	in int
}{
	{1},
	{2},
	{3},
	{10},
	{20},
}

func BenchmarkTmplExecute(b *testing.B) {
	b.ReportAllocs()
	temp := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		// 每次一次循环每一个goroutine都有自己单独的
		var buf bytes.Buffer
		for pb.Next() {
			// 这个循环每个goroutine都要循环b.NC
			buf.Reset()
			_ = temp.Execute(&buf, "World")
		}
	})
}
