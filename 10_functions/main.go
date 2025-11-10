package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// multiple return
func getLangs() (string, string, string, int) {
	return "gloang", "js", "python", 1
}

// callback func
func process(fn func(a int, b int) int) {
	fn(1, 2)
}

// variadic function : who accept multiple arguments
func variadicFunc(nums ...int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func main() {

	// sumOfNum := sum(1, 2)
	// fmt.Println(sumOfNum)

	lang1, lang2, lang3, num := getLangs()
	fmt.Println(lang1, lang2, lang3, num)

	process(sum)

	multiplesum := variadicFunc(1, 2, 3, 4, 5, 6)
	print(multiplesum)
}
