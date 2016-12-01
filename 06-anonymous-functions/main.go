package main

import "fmt"

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	/*Implement an inline lambda that squares items*/
	//items... is syntax for passing a slice(array) to a varargs function!
	squaredItems := applyOneEach(/* create a lambda that squares items*/, items...)
	/*Implement an inline lambda that cubes items*/
	cubedItems := applyOneEach(/* create a lambda that cubes items*/, items...)

	fmt.Println("Items", items)
	fmt.Println("Squares", squaredItems)
	fmt.Println("Cubes", cubedItems)
}

// This will be explained later... TODO maybe remove?
func applyOneEach(op func(int) int, items ...int) []int {
	var results []int
	for _, value := range items {
		results = append(results, op(value))
	}
	return results
}
