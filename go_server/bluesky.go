package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func getFeed() ([]FeedItem, error) {
	fmt.Println("Getting timeline...")

	getTimelineData := GetTimelineStruct{
		Limit: "5",
	}

	getTimelineBody, err := PrepJson(getTimelineData)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://bsky.social/xrpc/app.bsky.feed.getTimeline", getTimelineBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+Session.AccessJwt)
	req.Header.Add("Content-Type", "application/json")

	timelineResp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer timelineResp.Body.Close()

	timelineBody, err := io.ReadAll(timelineResp.Body)
	if err != nil {
		return nil, err
	}

	feed := FeedResponse{}
	if err := json.Unmarshal(timelineBody, &feed); err != nil {
		return nil, err
	}

	return feed.Feed, nil
}

func createSession() (SessionStruct, error) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	username := os.Getenv("USERNAME")
	if username == "" {
		return SessionStruct{}, fmt.Errorf("No username found!")
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		return SessionStruct{}, fmt.Errorf("No password found!")
	}

	// Log In and Create Session
	//
	fmt.Println("Creating session...")

	sessionLogin := LoginStruct{
		Identifier: username,
		Password:   password,
	}

	createSessionBody, err := PrepJson(sessionLogin)
	if err != nil {
		return SessionStruct{}, err
	}

	resp, err := http.Post(
		"https://bsky.social/xrpc/com.atproto.server.createSession",
		"application/json",
		createSessionBody)
	if err != nil {
		return SessionStruct{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	session := SessionStruct{}
	if err := json.Unmarshal(body, &session); err != nil {
		return SessionStruct{}, err
	}
	return session, nil
}

func makePost(session SessionStruct) {
	postReader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter post: ")
	postReader.Scan()
	postText := postReader.Text()

	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format(time.RFC3339)
	post := RecordStruct{
		CreatedAt: formattedTime,
		Text:      postText,
	}

	postInfo := PostStruct{
		Record:     post,
		Repo:       session.Handle,
		Collection: "app.bsky.feed.post",
	}

	postBody, err := PrepJson(postInfo)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://bsky.social/xrpc/com.atproto.repo.createRecord", postBody)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+session.AccessJwt)
	req.Header.Add("Content-Type", "application/json")

	postResp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer postResp.Body.Close()

	if postResp.StatusCode == 200 {
		fmt.Println("Post created :)")
	} else {
		fmt.Println("Error in posting! Sorry :(")
	}
}
