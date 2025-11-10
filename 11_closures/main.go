package main

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {

	increment := counter()

	print(increment()) //counter is 1
	print(increment()) //counter is 2
	//as the count is used by the function called inside counter so it does not removed even  the counter func removed from call-stack
	//as we know very well the variables values get removed from call-stack after the function execution

}
