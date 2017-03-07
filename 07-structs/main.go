package main

type Tweet struct {
	User    string
	Message string
	Likes   int
}

func main() {
	// Tweet created by listing fields in order
	okkesTweet := &Tweet{"Okke", "Hallo gophers", 12}
	// Tweet created with named parameters, missing fields will be initialized to their null values!
	chrisTweet := &Tweet{User: "Christophe", Message: "Lalala wat een fijne dag"}

	// Create a function Print() on Tweet that prints all the fields of tweet
	okkesTweet.Print()
	chrisTweet.Print()
}
