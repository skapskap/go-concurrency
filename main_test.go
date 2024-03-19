package main

import (
	"sync"
	"testing"
)

func BenchmarkSynchronousFetching(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syncUsername := fetchUserSync()
		syncLikes := fetchUserLikesSync(syncUsername)
		syncMatch := fetchUserMatchSync(syncUsername)

		_ = syncLikes
		_ = syncMatch
	}
}

func BenchmarkAsynchronousFetching(b *testing.B) {
	for i := 0; i < b.N; i++ {
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

		_ = asyncLikes
		_ = asyncMatch
	}
}
