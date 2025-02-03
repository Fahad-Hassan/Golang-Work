package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("heloow world")
	// time.Sleep(2000 * time.Microsecond)
	fmt.Println("endd")
}
func sayHi() {
	fmt.Println("hi fahad")
	time.Sleep(1000 * time.Microsecond)

}

func main() {
	fmt.Println("learning goroutines")
	go sayHello()
	go sayHi()
	time.Sleep(1000 * time.Microsecond)

}
