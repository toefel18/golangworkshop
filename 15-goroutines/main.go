package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    waitgroup := sync.WaitGroup{} //CountDownLatch in Java!
    waitgroup.Add(3)
    go printer(&waitgroup, []string{"a", "b", "c", "d", "e", "f", "g"})
    go printer(&waitgroup, []string{"I", "II", "III", "IV", "V", "VI", "VII"})

    // write a function intPrinter that prints 10 integers (one per second)
    // and run that in a separate goroutine!
    go intPrinter(&waitgroup, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

    waitgroup.Wait()
    fmt.Print("All goroutines finished")
}

func printer(waitgroup *sync.WaitGroup, slice []string) {
    for _, v := range slice {
        fmt.Print(v, " ")
        time.Sleep(1 * time.Second)
    }
    waitgroup.Done()
}

func intPrinter(waitgroup *sync.WaitGroup, slice []int) {
    for _, v := range slice {
        fmt.Print(v, " ")
        time.Sleep(1 * time.Second)
    }
    waitgroup.Done()
}