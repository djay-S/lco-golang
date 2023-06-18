package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var parallelCounter int = 0
var mutexedParallelCounter int = 0
var waitGroup sync.WaitGroup
var mutex sync.Mutex

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("\nCounter:", (i + 1))
		normalAdder()
		parallelizedAdder()
		mutexedParallelAdder()

		fmt.Println("Counter value is:", counter)
		fmt.Println("Parallel Counter value is:", parallelCounter)
		fmt.Println("Mutexed Parallel Counter value is:", mutexedParallelCounter)
		counter = 0
		parallelCounter = 0
		mutexedParallelCounter = 0
	}
}

func normalAdder() {
	normalAdd()
	normalAdd()
}

func normalAdd() {
	for i := 0; i < 10000000; i++ {
		counter++
	}
}

func parallelizedAdder() {
	go parallelizedAdd()
	waitGroup.Add(1)

	go parallelizedAdd()
	waitGroup.Add(1)

	waitGroup.Wait()
}

func parallelizedAdd() {
	defer waitGroup.Done()

	for i := 0; i < 10000000; i++ {
		parallelCounter++
	}
}

func mutexedParallelAdder() {
	go mutexedParallelizedAdd()
	waitGroup.Add(1)

	go mutexedParallelizedAdd()
	waitGroup.Add(1)

	waitGroup.Wait()
}

func mutexedParallelizedAdd() {
	defer waitGroup.Done()

	for i := 0; i < 10000000; i++ {
		mutex.Lock()
		mutexedParallelCounter++
		mutex.Unlock()
	}
}
