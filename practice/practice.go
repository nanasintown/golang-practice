package main

/**
1. Reversing an Array: Write a function that reverses an array of integers in place.

*/

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("array and slice coding questions")
	arrTest := []int{1, 4, 2, 3, 19, 39, 24, 16, 11, 87, 65, 29}
	// reverse(arrTest)
	// resultDup := findDup(arrTest)
	// fmt.Println(resultDup)
	// result := TwoSum(arrTest, 40)
	result := remove(arrTest)

	fmt.Println(result)
}

func reverse(arr []int) {
	first, last := 0, len(arr)-1
	for first < last {
		arr[first], arr[last] = arr[last], arr[first]
		first++
		last--
	}
}

/*
2. Find Duplicates: Given an array of integers, find and return the duplicates.
	 Implement it with optimal space and time complexity.
*/

func findDup(arr []int) []int {
	result := []int{}
	checkInt := make(map [int]int)

	for _, num := range arr {
		checkInt[num]++
		if checkInt[num] > 1 {
			result = append(result, num)
		}

	}
	return result
}

/*
3. Two Sum Problem: Given an array of integers, find two numbers such that they
	 add up to a specific target. Implement this using a hash map for optimal time complexity.
*/
func TwoSum(arr []int, target int) []int {
	sums := make(map[int]int)
	for i, num := range arr {
		diff := target - num
		fmt.Println(num, diff)
		// 	arrTest := []int{1, 4, 19, 39, 24, 16, 11, 87, 65, 29}
		if j, ok := sums[diff]; ok{
			return []int{j, i}
		}
		sums[num] = i 
	}
	return nil
} 

func remove(arr []int) int {
	if (len(arr) == 0){
		return 0
	}
	sort.Ints(arr)
	// slices.Sort(arr)

	var prev int = arr[0]
	var kept int = 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > prev + 1 {
			kept++
			prev = arr[i]
		}
	}
	return len(arr) - kept
}