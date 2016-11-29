package main

import (
	"fmt"
	"errors"
	"github.com/toefel18/cov-experiment/random"
)

type Tweet struct {
	User string
	Message string
	Likes int
}

func (t * Tweet) Print() {
	fmt.Printf("%v: %v  (likes: %v)\n", t.User, t.Message, t.Likes)
}

func main() {
	// Tweet created by listing fields in order
	okkesTweet := &Tweet{"Okke", "Hallo gophers", 12}
	// Tweet created with named parameters, missing fields will be initialized to their null values!
	chrisTweet := &Tweet{User: "Christophe", Message: "Lalala wat een fijne dag"}

	//Print the tweet with the most likes!
	if okkesTweet.Likes > chrisTweet.Likes {
		okkesTweet.Print()
	} else {
		chrisTweet.Print()
	}
}

func commonIfStatementUsage() {
	// if <statement> ; <condition> { }
	if randomNr, err := computeRandomNumber(); err != nil {
		// something bad happened, handle it NOW
	} else {
		fmt.Println(randomNr)
	}
	// randomNr is out of scope
}

func computeRandomNumber() (int, error) {
	return random.NextInt(), errors.New("failed to generate random number")
}

