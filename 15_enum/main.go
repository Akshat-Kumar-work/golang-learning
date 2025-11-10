package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status", status)
}

type ApplicationStatus string

const (
	Completed     ApplicationStatus = "completed"
	UnderProgress ApplicationStatus = "under_progress"
	Rejected      ApplicationStatus = "rejected"
)

func changeApplicationStatus(status ApplicationStatus) {
	fmt.Println("changing application status", status)
}

func main() {
	changeOrderStatus(Received)
	changeApplicationStatus(Rejected)
}
