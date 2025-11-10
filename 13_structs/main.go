package main

import (
	"fmt"
)

// struct are similar to class
type Person struct {
	name   string
	age    int
	gender string
}

// receiver type (value receiver) means struct are pass by value by default
// connecting function with our struct or we can say making this function as property of our struct
func (P Person) changeGender(gender string) {
	P.gender = gender
}

// pass by reference, we always need to use * (pointer) we have to modify the value
func (P *Person) changeGenderByReference(gender string) {
	P.gender = gender
}

// constructor for Person struct
func NewPerson(name string, age int, gender string) *Person {
	myPerson := Person{
		name:   name,
		age:    age,
		gender: gender,
	}
	return &myPerson //returning memory location or we can say instance of that person struct
}

// struct embeding or relations/inheritance
type Car struct {
	petrol int
	color  string
}

type Maruti struct {
	wheel int
	Car   //inheritance
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

	myperson := NewPerson("ok", 3, "female")
	fmt.Println(myperson)

	//other way to define and initialize struct (if we are going to use only once)
	language := struct {
		name string
		code string
	}{"hindi", "hin-IN"}

	fmt.Println(language)

	//struct embedings (relations or inheritance)
	marutiCar := Maruti{wheel: 4, Car: Car{color: "white", petrol: 3}}
	fmt.Println(marutiCar)
}
