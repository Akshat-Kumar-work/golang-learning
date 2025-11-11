package main

import (
	"fmt"
)

// T is a type parameter
// here T is generic, any / interface{}
func PrintSlice[T int | string](slice []T) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

// generic in struct
type stack[T any] struct {
	elements []T
}

func main() {
	ints := []int{1, 2, 3}
	strings := []string{"a", "b", "c"}

	PrintSlice(ints)    // T = int
	PrintSlice(strings) // T = string

	myStack := stack[string]{
		elements: []string{"gloang"},
	}
	fmt.Println(myStack)
}
