package main

/**
1. Reversing an Array: Write a function that reverses an array of integers in place.

*/

import "fmt"

func main() {
	fmt.Println("array and slice coding questions")
	arrTest := []int{1, 4, 19, 39, 24, 16, 11, 87, 65, 29}
	// reverse(arrTest)
	// result := findDup(arrTest)
	// fmt.Println(result)
	result := TwoSum(arrTest, 40)
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
		if j, ok := sums[diff]; ok{
			return []int{j, i}
		}
		sums[num] = i 
	}
	return nil
} 

