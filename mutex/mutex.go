package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() // Unlock after modification
		wg.Done()
	}()

	p.mu.Lock() // Lock before modification
	p.views += 1

}

// Avoiding race condition with mutex
func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for range 100 {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println("Views:", myPost.views)

}

// The below creates a race condition
// func (p *post) inc(wg *sync.WaitGroup) {
// 	defer func() {
// 		wg.Done()
// 	}()
// 	p.views += 1
// }
// func main() {
// 	var wg sync.WaitGroup
// 	myPost := post{views: 0}

// 	for range 100 {
// 		wg.Add(1)
// 		go myPost.inc(&wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Views:", myPost.views)

// }
