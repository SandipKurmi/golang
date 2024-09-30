package main

import (
	"fmt"
	"sync"
)

// goroutines
func task(id int, wg *sync.WaitGroup) {
	fmt.Println("task", id)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go task(i, &wg) // this will create 100000 go routines
	}

	wg.Wait()
}