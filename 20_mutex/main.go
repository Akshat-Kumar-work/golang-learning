package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

//Mutex stands for “mutual exclusion”.
//It’s a synchronization primitive that ensures only one goroutine at a time can access a shared piece of data.
//Basically it prevent from race conditions like multiple thread try to update the views parallely
//so by using mutex we lock the views field so that only 1 thread/goRoutine can update the field at a time

func (p *post) incrementViewOnPost(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views += 1
}

func main() {
	mypost := post{views: 0}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mypost.incrementViewOnPost(&wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final views:", mypost.views)
}
