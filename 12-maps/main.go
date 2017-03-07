package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["java"] = 6
	m["ruby"] = 9
	m["go"] = 8
	m["python"] = 5

	// Implement printScore, which takes a map[String]int and a string language and
	// prints the score of the language if the language exists, otherwise print
	// "language does not exist"

	// ranging over a map gives keys and values, if value or key is not used, then use a _
	for k, _ := range m {
		printScore(m, k)
	}

}
