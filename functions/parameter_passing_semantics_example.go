package main

import "fmt"

type Counter struct {
	val int
}

// reference semantics
func Increment(c *Counter) {
	c.val++
}

// changes are lost
func IncreaseByAmtWrong(c Counter, amt int) {
	c.val += amt
}

// changes
func IncreaseByAmt(c *Counter, amt int) {
	c.val += amt
}

func updateCounter(c *Counter) {
	// this works
	Increment(c)
	fmt.Println(c, "Expected", 1)

	// this doesn't
	IncreaseByAmtWrong(*c, 3)
	fmt.Println(c, "Expected", 4)

	// this works
	IncreaseByAmt(c, 3)
	fmt.Println(c, "Expected", 4)
}

func parameterPassingSemanticsExample() {
	c := Counter{}
	updateCounter(&c) // passing pointer
}
