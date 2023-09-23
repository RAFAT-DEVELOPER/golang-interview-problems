package main

import (
	"testing"
)

func BenchmarkConcatPlusOperator(b *testing.B) {
	a := "Hello, "
	benchmarkConcat(concatPlusOperator, a, b)
}

func BenchmarkConcatFmtSprintf(b *testing.B) {
	a := "Hello, "
	benchmarkConcat(concatFmtSprintf, a, b)
}

func BenchmarkConcatStringBuilder(b *testing.B) {
	a := "Hello, "
	benchmarkConcat(concatStringBuilder, a, b)
}

func benchmarkConcat(fn func(string, string) string, a string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fn(a, "World!")
	}
}
