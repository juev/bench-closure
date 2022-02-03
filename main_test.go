package main

import "testing"

func Benchmark_closure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		closure()
	}
}

func Benchmark_run(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run()
	}
}
