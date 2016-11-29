package main

import "fmt"

type Likeable interface {
	Like()
}

type Tweet struct {
	User string
	Message string
	Likes int
}

func (t *Tweet) Like() {
	t.Likes = t.Likes + 1
}

func main() {
	okkesTweet := &Tweet{"Okke", "Hallo gophers", 12}
	// Make Tweet implement the Likeable interface, the code below then compiles!
	var likableTweet Likeable = okkesTweet
	likableTweet.Like()
	fmt.Println("Okkes tweet is now liked", okkesTweet.Likes, "times")
}



