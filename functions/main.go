package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// multiple return
func getLangs() (string, string, string, int) {
	return "gloang", "js", "python", 1
}

func process(fn func(a int, b int) int) {
	fn(1, 2)
}

func main() {

	// sumOfNum := sum(1, 2)
	// fmt.Println(sumOfNum)

	lang1, lang2, lang3, num := getLangs()
	fmt.Println(lang1, lang2, lang3, num)

	process(sum)
}
