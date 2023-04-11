package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("race condition in go lang learning golang")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("first func ")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("second func ")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("third func ")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Print(score)
}
