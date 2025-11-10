package main

import "fmt"

func main() {
	// map is an data structure to store key value pair
	//note : if key doesnot exist in map it return zero value
	//maps are un-ordered data structure

	//creating map
	mp := map[string]int{"ok": 3}
	fmt.Println(mp)

	//creatig map using make keyword
	m := make(map[string]string)

	//to set
	m["k"] = "nice"
	m["age"] = "40"

	//to get value via key
	fmt.Println(m["k"])

	delete(m, "k")

	val, exist := m["age"]
	fmt.Println(val)

	fmt.Println(exist)

	fmt.Print(m)
}
