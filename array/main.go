package main

import "fmt"

func main() {
	var nums [4]int

	nums[0] = 1

	//array length
	print(len(nums))

	print(nums[0])

	fmt.Println(nums)

	arr := [3]int{1, 2, 3}

	dynamicSize := [...]int{1, 2, 3}

	fmt.Print(arr)
	fmt.Print(dynamicSize)
}
