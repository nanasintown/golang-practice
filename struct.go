//go:build struct
// +build struct

package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func createPerson(name string) *Person {
	p := Person{name: name}
	p.age = 22
	return &p
}

// Struct embedding
type base struct {
	num int
}

type container struct {
	base // container embeds base => methods of base struct are also methods of struct container
	str  string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func main() {
	fmt.Println("Study struct")
	/*
	 Struct group together variable with different types
	 The variables in a struct are called fields
	 Structs are used to create more complex data types, and are fundamental of OOP in Go

	 Struct is mutable.

	 If a pointer is pointing to a struct, it is automaticall dereferenced

	 p := person{"happy", 8}
	 sp := &p

	 sp.age == 8 TRUE
	*/

	dog := struct {
		name string
		age  int
	}{
		"Pax",
		2,
	}
	fmt.Println(dog)

	co := container{
		base: base{num: 1},
		str:  "something",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("Describe:", co.describe())
}
