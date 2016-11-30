package main

import (
	"fmt"
	"errors"
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

	// call divide and write as: <statement>; <condition>
	if result, err := divide(3, 0); err != nil {
		fmt.Println("Something went south: ", err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return a / b, nil
	}
}

