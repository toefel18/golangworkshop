package main

import "fmt"


func main() {
    fmt.Println("add     4 + 5 =", operation(add, 4, 5))
    fmt.Println("product 4 + 5 =", operation(product, 4, 5))
}

// Implement the function 'operation' that accepts a function and two ints
// return the result of invoking the function on a and b
func operation(op func(a, b int) int, a, b int) int {
    return op(a, b)
}

func add(a, b int) int {
    return a + b
}

func product(a, b int) int {
    return a * b
}
