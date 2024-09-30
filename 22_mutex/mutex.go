package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex // mutex
}


func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer func ()  { 
	 p.mu.Unlock() // unlock mutex
	 wg.Done()
	}()

	p.mu.Lock() // lock mutex
	p.views++
}

func main() {

	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			myPost.incrementViews(&wg)
		}()
	}

	wg.Wait()

	fmt.Println(myPost.views)
}