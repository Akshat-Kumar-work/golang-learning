package main

import "fmt"

const num = 1

func main() {
	const name = "akshat"
	const (
		port = 500
		host = "localhost"
	)
	print(num)
	fmt.Print(port, host)
}
