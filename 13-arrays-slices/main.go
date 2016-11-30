package main

import "fmt"

func main() {
    strArray := make([]string, 0)           // create a slice
    strArray = append(strArray, "Hi")       //append because an new backing array might be created to fit the new elem!
    strArray = append(strArray, "Gophers")
    strArray = append(strArray, "how")
    strArray = append(strArray, "are")
    strArray = append(strArray, "you")
    strArray = append(strArray, "liking")
    strArray = append(strArray, "it")

    printSlice(strArray)

    // take a slide of strArray and print that slice using a range for loop!
    firstTwo := strArray[:2]
    printSlice(firstTwo)

    // Print the length and capacity of firstTwo
    fmt.Println("length:", len(firstTwo), ", capacity:", cap(firstTwo))

    // append "of today" to firstTwo and print firstTwo again
    firstTwo = append(firstTwo, "of today")

    printSlice(firstTwo)

    //print strArray again, what do you expect to see?
    printSlice(strArray)

}

func printSlice(slice []string) {
    for _, v := range slice {
        fmt.Print(v, " ")
    }
    fmt.Println()
}