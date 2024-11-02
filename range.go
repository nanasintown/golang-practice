//go:build range
// +build range

package main

import "fmt"

func main() {
	nums := []int{4, 2313, 52389578} // make a slice of numbers

	// use range to sum the slice, do we have sum function()?
	// For array is the same

	sum := 0

	// range provide the index and value, here _ is the index, which is not needed now
	for _, n := range nums {
		sum += n
	}
	fmt.Println("Sum is: \n", sum)

	// use range on maps
	m1 := map[string]int{"happy": 7, "nana": 22}
	for k, v := range m1 {
		fmt.Printf("Name %s, age %d\n", k, v) // USE Printf not ln because of format (or string literal)
	}

	for k := range m1 {
		fmt.Println("Only key: ", k)
	}

	for _, v := range m1 {
		fmt.Printf("Only value %d\n", v)
	}

	// range on string iterate over Unicode code points

	for i, c := range "nana" {
		fmt.Println(i, c)
	}
}
