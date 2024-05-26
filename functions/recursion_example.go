package main

import "fmt"

func recursionExample() {
	var fib func(int) int // need this variable so that it can be used in the body of lambda

	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("fib(7):", fib(7))
}
