package main

func main() {
	age := 16

	if age >= 18 {
		print("greater than 18")
	} else if 15 < age && age < 18 {
		print("greather than 15 less than 18")
	} else {
		print("don't now")
	}

	//logical operators
	// == equal
	// or ||
	// and &&

	//we can also directly declare variables inside if and provide conditions
	if marks := 1; marks > 0 {
		print("marks is greater than 0")
	}
}
