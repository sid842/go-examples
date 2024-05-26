package main

import "fmt"

func intSeq() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

func closureExample() {
	nextInt := intSeq()

	fmt.Println("num:", nextInt())
	fmt.Println("num:", nextInt())
	fmt.Println("num:", nextInt())

	newInts := intSeq()
	fmt.Println("new seq:", newInts())
}
