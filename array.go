//go:build array
// +build array

package main

import "fmt"

func main() {
	// array has a fixed size. We can change the value of the items, but not the size of the array
	var a [3]int   // initialize an array with 3 items, 0 is default value if not specified
	fmt.Println(a) // a => [0 0 0]
	aa := [2]string{"ss", "oo"}
	fmt.Println(aa)


	b := [4]int{1, 4, 6, 2}
	fmt.Println(b)

	b = [...]int{1, 2, 3, 5}
	fmt.Println("dcl:", b)

	// 2 dimensional array
	var arr2 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr2[i][j] = i + j
		}
	}
	fmt.Println(arr2)

	var twoD [2][3]int
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
