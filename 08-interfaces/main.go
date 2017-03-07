package main

import "fmt"

type Likeable interface {
	Like()
}

type Tweet struct {
	User    string
	Message string
	Likes   int
}

func main() {
	okkesTweet := &Tweet{"Okke", "Hallo gophers", 12}
	// Make Tweet implement the Likeable interface, the code below should then compile and print 13!
	var likableTweet Likeable = okkesTweet
	likableTweet.Like()
	fmt.Println("Okkes tweet is now liked", okkesTweet.Likes, "times")
}
