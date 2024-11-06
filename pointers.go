//go:build pointers
// +build pointers

package main

import "fmt"

func main() {
	fmt.Println("Pointer in Go lang")
	num := 24
	// var numPtr *int
	// numPtr = &num
	// var numPtr *int = &num

	numPtr := &num


	fmt.Println("From num var", &num)
	fmt.Println("Value of num pointer", numPtr)
	fmt.Println("Dereference:", *numPtr)

	*numPtr = 45
	fmt.Println("After change the value", *numPtr) // from 22 to 45
	fmt.Println("Num", num)
	fmt.Println(&num, numPtr)

}

/**
A pointer is a type of variable that store the address of another variable
num := 22, integer num has value of 22, and is located somewhere with address x
int* numptr is an int pointer, which can store the address numPtrof num
===> var numptr *int = &num THE ONLY CORRECT WAY
We can not assign int* numptr to the value of num like: int* numptr = num

If we print out the value of numptr, it will be the address of num,
and if we want to get the number 22 with variable numptr, we use dereference: *numptr

we can change the value of the variable through pointer (as pointer stores address)
===> *numptr = 29 (dereference and change the value)
*/
