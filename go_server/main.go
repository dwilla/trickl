package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	Session     SessionStruct
	Feed        []FeedItem
	CurrentPost int
)

func main() {
	var err error
	Session, err = createSession()
	if err != nil {
		panic(err)
	}

	Feed, err = getFeed()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handlePost)

	fmt.Println("Server starting on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
