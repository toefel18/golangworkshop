package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitgroup := sync.WaitGroup{} //CountDownLatch in Java!
	waitgroup.Add(2)
	go printer(&waitgroup, []string{"a", "b", "c", "d", "e", "f", "g"})
	go printer(&waitgroup, []string{"I", "II", "III", "IV", "V", "VI", "VII"})

	// write a function intPrinter that prints 10 integers (one per second)
	// and run that in a separate goroutine!

	waitgroup.Wait()
	fmt.Print("\nAll goroutines finished")
}

func printer(waitgroup *sync.WaitGroup, slice []string) {
	for _, v := range slice {
		fmt.Print(v, " ")
		time.Sleep(1 * time.Second)
	}
	waitgroup.Done()
}
