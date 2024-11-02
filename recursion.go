//go:build recursion
// +build recursion

package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return fact(n-1) * n
}

func main() {
	fmt.Println(fact(4))

	// closure as recursion
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(12))
}
