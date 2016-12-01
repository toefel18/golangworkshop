package main

import (
	"net/http"
	"encoding/json"
)

var (
	tweetsById = map[string]Tweet {
		"1" : Tweet{"Okke", "After all, all he did was string together a lot of old, well-known quotations.", 4, 2},
		"2" : Tweet{"Chris", "How many Zen masters does it take to screw in a light bulb?", 55, 89},
		"3" : Tweet{"Chef", "Your lucky color has faded.", 47, 667},
		"4" : Tweet{"Zanouzzi", "Courage, man; the hurt cannot be much.", 123, 2},
		"5" : Tweet{"Elle", "You will gain money by an illegal action.", 695, 2},
	}
)

type Tweet struct {
	User string
	Message string
	Likes int
	Retweets int
}

type Version struct {
	Name string
	Version string
}

func main() {
	// Write a simple HTTP server that publishes a simple Json REST API of tweets
	// when visiting http://localhost:8080/1  it should return okke's tweet in JSON
	// Hints:
	// 1. use http.NewServeMux()
	// 2. write a function and take a slice on the uri to get the ID
	mux := http.NewServeMux()
	mux.HandleFunc("/version/", versionHandler)
	mux.HandleFunc("/tweets/", tweetHandler)
	http.ListenAndServe(":8080", mux)
}

func versionHandler(resp http.ResponseWriter, req *http.Request) {
	if response, err := json.Marshal(&Version{"RestServer", "1.2.3"}); err != nil {
		resp.WriteHeader(500)
	} else {
		resp.Write(response)
	}
}

func tweetHandler(resp http.ResponseWriter, req *http.Request) {
	id := req.RequestURI[8:]
	if tweet, exists := tweetsById[id]; exists {
		if response, err := json.Marshal(&tweet); err != nil {
			resp.WriteHeader(500)
		} else {
			resp.Write(response)
		}
	} else {
		resp.WriteHeader(404)
	}

}