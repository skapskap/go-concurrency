package main

import (
	"fmt"
	"sync"
	"time"
)

// Synchronous Fetching - Result: 200.56ms

func main()  {
	start := time.Now()
	username := fetchUser()
	respCh := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fetchUserLikes(username, respCh, wg)
	go fetchUserMatch(username, respCh, wg)

	wg.Wait() // Block the execution until 2 wg.Done() are called
	close(respCh) // Channel closing before the goroutines finish their task - use waitgroup

	for resp := range respCh {
		fmt.Println("resp: ", resp)
	}

	// fmt.Println("likes: ", likes)
	// fmt.Println("match: ", match)
	fmt.Println("took: ", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "Lucas"
}

func fetchUserLikes(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respCh <- 11
	wg.Done()
}

func fetchUserMatch(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respCh <- "Sabrina"
	wg.Done()
}

// Synchronous Fetching - Result: 300.64ms

// func main()  {
// 	start := time.Now()
// 	username := fetchUser()
// 	likes := fetchUserLikes(username)
// 	match := fetchUserMatch(username)

// 	fmt.Println("likes: ", likes)
// 	fmt.Println("match: ", match)
// 	fmt.Println("took: ", time.Since(start))
// }

// func fetchUser() string {
// 	time.Sleep(time.Millisecond * 100)

// 	return "Lucas"
// }

// func fetchUserLikes(username string) int {
// 	time.Sleep(time.Millisecond * 100)

// 	return 11
// }

// func fetchUserMatch(username string) string {
// 	time.Sleep(time.Millisecond * 100)

// 	return "Sabrina"
// }