package main

import (
	"fmt"
)

// Approach 1: Brute Force (Time: O(n^2), Space: O(1))

func findZeroSumSubarraysBruteForce(arr []int) [][]int {
	n := len(arr)
	subarrays := [][]int{}

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == 0 {
				subarrays = append(subarrays, arr[i:j+1])
			}
		}
	}

	return subarrays
}

// Approach 2: Prefix Sum (Time: O(n), Space: O(n))

func findZeroSumSubarraysPrefixSum(arr []int) [][]int {
	n := len(arr)
	subarrays := [][]int{}
	prefixSum := make([]int, n+1)
	prefixSum[0] = 0
	sumToIndices := make(map[int][]int)

	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i-1]
		if prefixSum[i] == 0 {
			subarrays = append(subarrays, arr[0:i])
		}
		if indices, found := sumToIndices[prefixSum[i]]; found {
			for _, index := range indices {
				subarrays = append(subarrays, arr[index+1:i])
			}
		}
		sumToIndices[prefixSum[i]] = append(sumToIndices[prefixSum[i]], i-1)
	}

	return subarrays
}

// Approach 3: Hash Table (Time: O(n), Space: O(n))

func findZeroSumSubarraysHashTable(arr []int) [][]int {
	n := len(arr)
	subarrays := [][]int{}
	sumToIndices := make(map[int][]int)
	sum := 0

	for i := 0; i < n; i++ {
		sum += arr[i]
		if sum == 0 {
			subarrays = append(subarrays, arr[0:i+1])
		}
		if indices, found := sumToIndices[sum]; found {
			for _, index := range indices {
				subarrays = append(subarrays, arr[index+1:i+1])
			}
		}
		sumToIndices[sum] = append(sumToIndices[sum], i)
	}

	return subarrays
}

// Approach 4: Kadane's Algorithm (Time: O(n), Space: O(1))

func findZeroSumSubarraysKadane(arr []int) [][]int {
	n := len(arr)
	subarrays := [][]int{}
	sumToIndices := make(map[int][]int)
	currentSum := 0
	startIndex := -1

	for i := 0; i < n; i++ {
		currentSum += arr[i]

		if currentSum == 0 {
			subarrays = append(subarrays, arr[startIndex+1:i+1])
		}

		if indices, found := sumToIndices[currentSum]; found {
			for _, index := range indices {
				subarrays = append(subarrays, arr[index+1:i+1])
			}
		} else {
			sumToIndices[currentSum] = append(sumToIndices[currentSum], i)
		}
	}

	return subarrays
}

func main() {
	fmt.Println("The Grand Zero-Sum Subarray Adventure Begins! 🌟🚀")

	input1 := []int{2, 3, -1, 2, -4}
	fmt.Println("Input 1:")

	fmt.Println("Brute Force:")
	subarrays1BF := findZeroSumSubarraysBruteForce(input1)
	fmt.Println(subarrays1BF)

	fmt.Println("Prefix Sum:")
	subarrays1PS := findZeroSumSubarraysPrefixSum(input1)
	fmt.Println(subarrays1PS)

	fmt.Println("Hash Table:")
	subarrays1HT := findZeroSumSubarraysHashTable(input1)
	fmt.Println(subarrays1HT)

	fmt.Println("Kadane's Algorithm:")
	subarrays1KA := findZeroSumSubarraysKadane(input1)
	fmt.Println(subarrays1KA)

	input2 := []int{-1, 3, 1, -3, 1, 5, -4, 1, 3}
	fmt.Println("Input 2:")
	fmt.Println("Brute Force:")

	subarrays2BF := findZeroSumSubarraysBruteForce(input2)

	fmt.Println(subarrays2BF)
	fmt.Println("Prefix Sum:")

	subarrays2PS := findZeroSumSubarraysPrefixSum(input2)
	fmt.Println(subarrays2PS)
	fmt.Println("Hash Table:")

	subarrays2HT := findZeroSumSubarraysHashTable(input2)

	fmt.Println(subarrays2HT)

	fmt.Println("Kadane's Algorithm:")
	subarrays2KA := findZeroSumSubarraysKadane(input2)
	fmt.Println(subarrays2KA)

	fmt.Println("The Grand Zero-Sum Subarray Adventure Ends! 🎉🚀")
}
