package main

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

func main() {
	BenchMarkResultExample()
}

func BenchMarkResultExample() {
	// 基准测试测试并行并发
	benchMarkResult := testing.Benchmark(func(b *testing.B) {
		temp := template.Must(template.New("test").Parse("Hello,{{.}}"))
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// Each goroutine has its own bytes.Buffer.
			var buf bytes.Buffer
			for pb.Next() {
				// The loop body is executed b.N times total across all goroutines.
				buf.Reset()
				_ = temp.Execute(&buf, "world")
			}
		})
	})

	fmt.Printf("%s\t%s\n", benchMarkResult.String(), benchMarkResult.MemString())
}
