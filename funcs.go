//go:build funcs
// +build funcs

package main

import "fmt"

// To declare a function, pass in the type of parameter if needed
// then also define the type of return value
func plusPlus(a, b, c int) int {
	return a + b + c
}

// multiple return values then put it in ()
func vals() (int, int) {
	return 22, 21
}

// Variadic functions take flexible number of arguments, like the Println function

func sum(num ...int) {
	fmt.Print(num, " ")
	total := 0

	for _, n := range num {
		total += n
	}

	fmt.Println("total", total)
}

// Closure and Anonymous functions
// A closure in Go is a function value that references variables from outside its body.
func intSeq() func() int {
	i := 0 // this variable is not passed to the closure, but the closure can still access it
	return func() int {
		i++
		return i
	}
}

/*
Function ở trên tên là outer, định nghĩa và trả
về một function khác nhằm tăng và trả về
biến x được định nghĩa trong phạm vi bên ngoài.
 Bạn có thể sử dụng function bên trong này như sau:
 increment := intSeq() // tạo tên gọi function là increment

increment() // 1
increment() // 2
increment() // 3

*/

func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(x int) int {
		if y, ok := cache[x]; ok { // ok is a special keyword.
			return y
		}
		y := f(x) // == fib(x)
		cache[x] = y
		return y
	}
}

// memoize returns an anonymous function, and that anonymous function needs an int as argument
// the anonymous function then returns an int
// so if var1 := memoize, then var1 is a function expecting for an int and will return an int
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(vals()) // Get all by once

	a, b := vals() // a = 22, b = 21
	fmt.Println(a)
	fmt.Println(b)

	res := plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...)
	nums := []int{12, 31242, 52, 7567}
	sum(nums...)

	// can also pass the arguments like this
	sum(3231, 8979)

	fastFib := memoize(fib)  // fastFib is now a function that is expecting an integer as argument
	fmt.Println(fastFib(10)) // Calculate Fibonacci number efficiently

}
