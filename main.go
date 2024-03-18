package main

import (
	"fmt"
	"time"
)

// Synchronous Fetching - Result: 300.64ms

func main()  {
	start := time.Now()
	username := fetchUser()
	likes := fetchUserLikes(username)
	match := fetchUserMatch(username)

	fmt.Println("likes: ", likes)
	fmt.Println("match: ", match)
	fmt.Println("took: ", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "Lucas"
}

func fetchUserLikes(username string) int {
	time.Sleep(time.Millisecond * 100)

	return 11
}

func fetchUserMatch(username string) string {
	time.Sleep(time.Millisecond * 100)

	return "Sabrina"
}