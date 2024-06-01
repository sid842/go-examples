package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) printPersonalDetails() {
	fmt.Println("Name:", p.name, "Age:", p.age)
}

type employee struct {
	person
	compnay    string
	experience int
}

func (e employee) printWorkExperience() {
	fmt.Printf("%s has been working in %s for %d years\n", e.name, e.compnay, e.experience)
}

func structEmbeddingExample() {
	emp := employee{
		person: person{
			name: "sid",
			age:  29,
		},
		compnay:    "meta",
		experience: 3,
	}

	emp.printPersonalDetails()
	emp.printWorkExperience()
}
