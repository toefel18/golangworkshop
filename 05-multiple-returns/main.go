package main

import "fmt"

func main() {
    a, b := "Go", "Oh I like"
    x, y := swap(a, b)
    fmt.Println(x, y)
}

// Implement a function function that takes two strings and returns them swapped
func swap(a, b string) (string, string) {
    return b, a
}