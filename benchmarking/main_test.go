package main

import "testing"

func BenchmarkTwoSumWithBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSumWithBruteForce([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSumWithTwoPassHashTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSumWithTwoPassHashTable([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSumOnePassHashTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSumOnePassHashTable([]int{2, 7, 11, 15}, 9)
	}
}
