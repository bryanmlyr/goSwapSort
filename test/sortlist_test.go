package main

import (
	"testing"

	sp "main/internal/swapSort"
)

func BenchmarkSortList10000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		sp.SortList(10000, false)
	}
}

func BenchmarkSortList1000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		sp.SortList(1000, false)
	}
}

func BenchmarkSortList100(b *testing.B) {
	for x := 0; x < b.N; x++ {
		sp.SortList(100, false)
	}
}

func BenchmarkSortList10(b *testing.B) {
	for x := 0; x < b.N; x++ {
		sp.SortList(10, false)
	}
}
