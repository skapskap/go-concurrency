package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Synchronous Fetching - Result: 300.69ms
	startSync := time.Now()
	syncUsername := fetchUserSync()
	syncLikes := fetchUserLikesSync(syncUsername)
	syncMatch := fetchUserMatchSync(syncUsername)

	fmt.Println("Synchronous Fetching:")
	fmt.Println("Username:", syncUsername)
	fmt.Println("Likes:", syncLikes)
	fmt.Println("Match:", syncMatch)
	fmt.Println("Synchronous Fetching took:", time.Since(startSync))

	fmt.Println()

	// Asynchronous Fetching - Result: 200.42ms
	startAsync := time.Now()
	asyncUsername := fetchUserAsync()
	respCh := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fetchUserLikesAsync(asyncUsername, respCh, wg)
	go fetchUserMatchAsync(asyncUsername, respCh, wg)

	wg.Wait()
	close(respCh)

	var asyncLikes int
	var asyncMatch string
	for resp := range respCh {
		switch v := resp.(type) {
		case int:
			asyncLikes = v
		case string:
			asyncMatch = v
		}
	}

	fmt.Println("Asynchronous Fetching:")
	fmt.Println("Username:", asyncUsername)
	fmt.Println("Likes:", asyncLikes)
	fmt.Println("Match:", asyncMatch)
	fmt.Println("Asynchronous Fetching took:", time.Since(startAsync))
}

func fetchUserSync() string {
	time.Sleep(time.Millisecond * 100)
	return "Lucas"
}

func fetchUserLikesSync(username string) int {
	time.Sleep(time.Millisecond * 100)
	return 11
}

func fetchUserMatchSync(username string) string {
	time.Sleep(time.Millisecond * 100)
	return "Sabrina"
}

func fetchUserAsync() string {
	time.Sleep(time.Millisecond * 100)
	return "Lucas"
}

func fetchUserLikesAsync(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respCh <- 11
	wg.Done()
}

func fetchUserMatchAsync(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respCh <- "Sabrina"
	wg.Done()
}
