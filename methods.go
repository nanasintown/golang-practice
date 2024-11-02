//go:build methods
// +build methods

package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Method to modify the dimensions of the rectangle
func (r *rect) setDimensions(width, height int) {
	r.width = width
	r.height = height
}

// Method has a receiver defined in front of the method's name. The receiver will receive the method
// so the receiver can call the method directly, like method in a class

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Initial area: ", r.area())
	fmt.Println("Initial perim:", r.perim())

	// Modify the dimensions of the rectangle
	r.setDimensions(20, 10)

	fmt.Println("Modified area: ", r.area())
	fmt.Println("Modified perim:", r.perim())

	rp := &r
	fmt.Println("Pointer area: ", rp.area())
	fmt.Println("Pointer perim:", rp.perim())

	// Further modify the dimensions of the rectangle through the pointer
	rp.setDimensions(15, 7)

	fmt.Println("Further modified area: ", rp.area())
	fmt.Println("Further modified perim:", rp.perim())
}
