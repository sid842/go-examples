package main

import (
	"fmt"
	"maps"
)

func mapsExample() {
	m := make(map[string]int)

	m["k1"] = 4
	m["k2"] = 8

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3) // default zero value

	fmt.Println("len:", len(m))

	delete(m, "k2")
	clear(m)
	fmt.Println("map:", m)

	_, ok := m["k2"]
	fmt.Println("ok", ok)

	var nilMap map[string]int
	_, ok = nilMap["k1"] // okay to check nil map
	//nilMap["k4"] = 3      This will cause panic

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)

	// n2 := n
	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n2:", n2)

	if maps.Equal(n, n2) { // Equality of maps
		fmt.Println("n == n2")
	} else {
		fmt.Println("n != n2")
	}

	n2["foo"] = 4 // if n2 := n, then changes reflected in n as well
	fmt.Println("n:", n)

}
