package main

// Factorial contains an error, run the unit test and fix the implementation
// To execute the test: open a terminal and cd inside the 11-testing directory, then type:
// go test

func factorial(n int) int {
	if n == 0 {
		return 0
	} else {
		return recursiveFactorial(n, 1)
	}
}

func recursiveFactorial(n, current int) int {
	if n <= 0 {
		return current
	} else {
		return recursiveFactorial(n-1, current*n)
	}
}
