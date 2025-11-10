package main

import (
	"fmt"
	"slices"
)

func main() {
	//slice is basically an dyanmic array

	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:4] //slice from index 1 upto 4 (but index 4 is not included)
	fmt.Println(s)

	//empty slice
	again := []int{}
	print(again)

	//we can create a slice without an array using make keyword
	//make arguments type,inital length,capacity
	sl := make([]int, 2, 5)
	sl = append(sl, 2)
	sl[0] = 3
	fmt.Println(sl)

	//to copy slice
	ok := make([]int, len(sl))
	copy(ok, sl)
	fmt.Println("copied slice", ok)

	//slice operator
	sl2 := []int{1, 2, 3}

	fmt.Print(sl2[1:3])

	//some inbuilt slice methods
	fmt.Println(slices.Equal(sl, sl2))
}
