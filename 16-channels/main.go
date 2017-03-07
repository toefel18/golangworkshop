package main

import (
	"fmt"
	"time"
)

func main() {
	doneSignal := make(chan int)
	intStream := make(chan int)

	// Create two go-routines:
	// 1. The first go-routine produces an integer on the intStream channel every second
	// 2. The second go-routine consumes the produced integers and prints them, the second goroutine also holds
	//    a reference to the done channel
	// 3. When the first goroutine produced 10 integers, it closes the channel close(intStream)
	// 4. the second go-routine detects this, and sends a random int to the done channel, closing the program.
	// 5. Use no locks!

	<-doneSignal
	fmt.Println("Done!")
}
