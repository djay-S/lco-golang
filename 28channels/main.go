package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	waitGroup := &sync.WaitGroup{}

	myChannel := make(chan int, 2)

	// myChannel <- 5
	// fmt.Println(<-myChannel)

	waitGroup.Add(2)
	go func(myChannel chan int, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()

		fmt.Println(<-myChannel)
	}(myChannel, waitGroup)

	go func(myChannel chan int, waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()

		myChannel <- 5
		myChannel <- 6

		close(myChannel)
	}(myChannel, waitGroup)

	waitGroup.Wait()
}
