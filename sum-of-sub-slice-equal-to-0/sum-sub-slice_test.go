package main

import (
	"reflect"
	"testing"
)

func TestFindZeroSumSubarraysBruteForce(t *testing.T) {
	input := []int{-1, 3, 1, -3, 1, 5, -4, 1, 3}
	expected := [][]int{{-1, 3, 1, -3}, {1, -3, 1, 5, -4}, {-3, 1, 5, -4, 1}, {-4, 1, 3}}
	result := findZeroSumSubarraysBruteForce(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindZeroSumSubarraysPrefixSum(t *testing.T) {
	input := []int{-1, 3, 1, -3, 1, 5, -4, 1, 3}
	expected := [][]int{{-1, 3, 1, -3}, {1, -3, 1, 5, -4}, {-3, 1, 5, -4, 1}, {-4, 1, 3}}
	result := findZeroSumSubarraysPrefixSum(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindZeroSumSubarraysHashTable(t *testing.T) {
	input := []int{-1, 3, 1, -3, 1, 5, -4, 1, 3}
	expected := [][]int{{-1, 3, 1, -3}, {1, -3, 1, 5, -4}, {-3, 1, 5, -4, 1}, {-4, 1, 3}}
	result := findZeroSumSubarraysHashTable(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindZeroSumSubarraysKadanes(t *testing.T) {
	input := []int{-1, 3, 1, -3, 1, 5, -4, 1, 3}
	expected := [][]int{{-1, 3, 1, -3}, {1, -3, 1, 5, -4}, {-3, 1, 5, -4, 1}, {-4, 1, 3}}
	result := findZeroSumSubarraysKadane(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func BenchmarkFindZeroSumSubarraysBruteForce(b *testing.B) {
	input := []int{-1, 3, 1, -3, 1, 5, -4, 1, 3}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findZeroSumSubarraysBruteForce(input)
	}
}

func BenchmarkFindZeroSumSubarraysPrefixSum(b *testing.B) {
	input := []int{-1, 3, 1, -3, 1, 5, -4, 1, 3}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findZeroSumSubarraysPrefixSum(input)
	}
}

func BenchmarkFindZeroSumSubarraysHashTable(b *testing.B) {
	input := []int{-1, 3, 1, -3, 1, 5, -4, 1, 3}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findZeroSumSubarraysHashTable(input)
	}
}

func BenchmarkFindZeroSumSubarraysKadanesAlgorithm(b *testing.B) {
	input := []int{-1, 3, 1, -3, 1, 5, -4, 1, 3}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findZeroSumSubarraysKadane(input)
	}
}
