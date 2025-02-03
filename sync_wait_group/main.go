package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d statrted\n", i)
	//some task is
	fmt.Printf("worker %d end\n", i)

}
func main() {
	// fmt.Println("explore learning goroutines")

	//start 3 workers goroutine
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) //increment waitgroup counter
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("worker task completed")

}
