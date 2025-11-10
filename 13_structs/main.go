package main

import "fmt"

// struct are similar to class
type Person struct {
	name   string
	age    int
	gender string
}

// receiver type (value receiver) means struct are pass by value by default
func (P Person) changeGender(gender string) {
	P.gender = gender
}

// pass by reference
func (P *Person) changeGenderByReference(gender string) {
	P.gender = gender
}

func main() {
	// Create a struct value
	p1 := Person{name: "Akshat", age: 23}

	// Access fields
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	p1.changeGender("male")
	fmt.Println("changed gender", p1.gender)

	p1.changeGenderByReference("female")
	fmt.Println("changed gender by pass by reference", p1.gender)

	// Modify fields
	p1.age = 24
	fmt.Println("Updated age:", p1.age)
}
