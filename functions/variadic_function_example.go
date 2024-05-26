package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, n := range nums {
		total += n
	}

	fmt.Println(total)
}

func variadicFunctionExample() {
	sum(1, 2)
	sum(100, 200, 300, 600)

	nums := []int{1, 2, 3, 4, 5, 6}
	sum(nums...)
}
