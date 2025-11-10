package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		//println(nums[i])
	}

	// for index := range len(nums) {
	// 	print(nums[index])
	// }

	for index, value := range nums {
		println(index, value)
	}

	m := map[string]string{"name": "akshat", "age": "23"}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
