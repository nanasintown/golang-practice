//go:build slice
// +build slice

package main

import (
	"fmt"
)

func main() {
	// Slice is quite similar to array. It doesn't have fixed size, and we can change the value of the items

	var s []string
	fmt.Println("empty: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("empty 2: ", s, "len: ", len(s), "cap: ", cap(s)) // cap = capacity

	// [x:y] the range includes x but exclude y

	// multi dimensional slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// append to slices. Creating slice in this way if you know the initial values of the elements
	// and the length of the slice, which may not change in the future. If potentially changed, use make
	s2 := []int{2, 3}
	fmt.Println(s2)
	s2 = append(s2, 4, 5)
	fmt.Println("s2 after append", s2)

	// create slice with make, this gives optional capacity, the elements are default set to 0
	s3 := make([]string, 2)
	s3[0] = "hihi"
	s3[1] = "hehe"
	fmt.Println(s3)
}
