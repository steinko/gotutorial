package main

import "testing"

func BenchmarkRepeat( b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeate("a",4)
	}
}