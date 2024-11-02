//go:build rune
// +build rune

package main

import (
	"fmt"
)

func main() {
	fmt.Println("In Go, strings are made of rune (character)")

	// const s = "สวัสดี"
	// fmt.Println("Len:", len(s))

	// // indexing a string accesses individual bytes, not characters
	// const s2 = "สวัสดี cao my nhut"
	// for i := 0; i < len(s2); i++ {
	// 	fmt.Printf("%x ", s2[i])
	// }
	// fmt.Println()
	// fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// var r rune = 'A'
	// fmt.Printf("Rune: %c, Unicode Code Point: %U, Integer Value: %d\n", r, r, r)

	// // Iterating over a string and printing runes
	// str := "Hello, 世界"
	// for index, r := range str {
	// 	fmt.Printf("Index: %d, Rune: %c, Unicode Code Point: %U, Integer Value: %d\n", index, r, r, r)
	// }
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%x ", str[i])
	// }

	var char rune = '♥'
	fmt.Println(char) // In ra ký tự hình trái tim

	// Chuyển đổi một chuỗi thành một slice các rune
	str2 := "世界!"
	runes := []rune(str2)
	fmt.Println(runes)

	fmt.Println(110 / 8)
}
