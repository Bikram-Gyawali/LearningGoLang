package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in go lang learning golang")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}
	// mut :=

	// myCh <- 5

	// fmt.Println(<-myCh)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen,val)
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		close(myCh)
		myCh <- 6
		myCh <- 3
		myCh <- 1
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
