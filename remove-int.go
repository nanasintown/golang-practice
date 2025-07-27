package main

import (
	// "sort"
	"slices"
)

func remove(arr []int) int {
	if (len(arr) == 0){
		return 0
	}
	// sort.Ints(arr)
	slices.Sort(arr)

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