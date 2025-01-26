package main

import (
	"fmt"
	"strconv"
)

func main(){
	var num int =42

	fmt.Printf("%T",num)
	var d float64 = float64(num)
	d =d+1.23
	fmt.Printf("%T",d)

	str :=strconv.Itoa(num)
	fmt.Printf("%T",str)
	
}