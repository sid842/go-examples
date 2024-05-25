package main

import (
	"fmt"
	"slices"
)

func slicesExample() {
	s := make([]string, 3)
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s[0], s[1], s[2] = "a", "b", "c"
	s = append(s, "d")
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	r := make([]string, len(s))
	copy(r, s)
	fmt.Println("r:", r, "len:", len(r), "cap:", cap(r))

	if slices.Equal(r, s) {
		fmt.Println("r == s")
	}

	v := s[1:3] // slice
	fmt.Println("v:", v, "len:", len(v), "cap:", cap(v))

	fmt.Println("appending in v, changes reflected in s")
	v = append(v, "e")

	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))
	fmt.Println("v:", v, "len:", len(v), "cap:", cap(v))

}
