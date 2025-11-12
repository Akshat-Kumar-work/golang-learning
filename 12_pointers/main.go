package main

import "fmt"

//pointers are the address of memory location of variables which we store
//map,slice and channels we don't need to pass the pointers
//for others we need to make them pass by reference

func changeNumbyvalue(num int) {
	num = 5
	fmt.Println("in changenum by value", num)
}

func changeNumByReference(num *int) {
	*num = 5
	fmt.Println("in changenum by reference ", num)
}

func main() {

	num := 1

	fmt.Println("memory address is", &num)

	//here we are passing by value
	changeNumbyvalue(num)

	fmt.Println("after change num by value", num)

	changeNumByReference(&num)

	fmt.Println("after change num by reference", num)

}
