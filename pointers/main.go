package main

import "fmt"

func main() {
	// var num int
	// num =2
	// var ptr *int
	// ptr = &num
	num := 2

	ptr := &num
	fmt.Println("address", ptr)
	fmt.Println("value", num)
	fmt.Println("ptr data",*ptr)
	var pointer *int

	fmt.Println(pointer)

	value :=10
	modifyvaluebyref(&value)
	fmt.Println(value)
}

func modifyvaluebyref(num *int){
	*num = *num*2
}

