package main

import "fmt"

//interfaces are contract
//we use er at end of interface as golang interface convention
//Interfaces describe behavior, not data.

//Implementation is implicit.

//A variable of interface type can hold any value that implements those methods.

// Define an interface
type Animal interface {
	MakeSound() string
}

// Define a type
type Dog struct{}

// Implement the interface method
// note: to implement interface method keep the method signature /name same
func (d Dog) MakeSound() string {
	return "Bark"
}

type Cat struct{}

func (c Cat) MakeSound() string {
	return "Meow"
}

func main() {
	var a Animal // interface variable

	a = Dog{}                  // Dog implements Animal
	fmt.Println(a.MakeSound()) // Bark

	a = Cat{}                  // Cat implements Animal
	fmt.Println(a.MakeSound()) // Meow
}
