package main

import "fmt"

func main() {
	you := yourName("Gopher, because I write GO!")
	fmt.Println("Hello", you())
}

// Implement a function yourName that takes a string
// return a function that returns a string when it's invoked
