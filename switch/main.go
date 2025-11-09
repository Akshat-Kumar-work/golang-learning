package main

import (
	"time"
)

func main() {
	i := 5

	switch i {
	case 1:
		print("i is 1")
	case 5:
		print("i is 5")
	}

	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		print("today is sat or sun so it is weekend")
	default:
		print("its work day")
	}
}
