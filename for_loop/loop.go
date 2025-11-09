package main

func main() {
	//while loop
	i := 1
	for i <= 3 {
		print(i)
		i += 1
	}

	//for loop
	for i := 0; i < 3; i++ {
		println(i)
	}

	//using range
	for j := range 3 {
		println(j)
	}

}
