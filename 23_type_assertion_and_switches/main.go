package main

import "fmt"

func main() {

	//type assertion is direct checking that an interface value holds specific type
	//it is an type checking
	var i any = "hello"
	value, ok := i.(string)
	if ok {
		fmt.Print("interface is type of string", value)
	}

	//type switch
	switch v := i.(type) {
	case int:
		fmt.Print("var is int")
	case string:
		fmt.Print("it is str")
	default:
		fmt.Print("unknown", v)
	}
}
