package main

import "fmt"

//enumerated types

// custom type
// type MyType string
type OrderStatus int
//for string type assign value to const 
const (
	Received OrderStatus = iota
	Confirm
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to ", status)
}

func main() {
	changeOrderStatus(Confirm)
}
