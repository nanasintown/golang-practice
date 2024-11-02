//go:build map
// +build map

package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("Make a map: map[key_type]<value_type>{key: value}")
	m := make(map[string]int)

	m["k1"] = 2002
	m["k2"] = 2017

	fmt.Println("Created map:", m)

	// Map is quite similar to dict in Python. Get the value by using the keys\

	v1 := m["k1"]
	fmt.Println(v1, "len of map", len(m))

	// remove a pair. If want to remove all, use clear(<map name>)
	delete(m, "k1")
	fmt.Println("Map m after delete first pair: ", m)

	unk, prs := m["k2"] // unk = value, prs = boolean, true if the key is present in the map
	fmt.Println(unk, prs)

	// declare and initialize with value
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// something fun from map package
	n2 := map[string]int{"a": 12, "b": 13}
	if maps.Equal(n, n2) {
		fmt.Println("Same map")
	} else {
		fmt.Println("Different map")
	}

	// can we sort the map? We can't sort it directly. An alternative way is to store the keys in a slice
	// sort that slice with package sort, then use it as a guid
	myMap := map[string]int{"c": 3, "a": 1, "b": 2}
	fmt.Print(myMap, "\n")
}
