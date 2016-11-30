package main

import "fmt"

func main() {
    defer fmt.Println("Goodbye")
   // write a defer statement that prints "Goodbye"
    fmt.Print("Hello ")
}