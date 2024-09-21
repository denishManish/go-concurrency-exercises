//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import (
	"testing"
	"time"
	"sync"
)

func TestMain(t *testing.T) {
	fetchSig := fetchSignalInstance()

	fail := make(chan bool)
	var wg sync.WaitGroup
	start := time.Unix(0, 0)
	wg.Add(1)
	go func(start time.Time) {
		defer wg.Done()
		defer close(fail)
		for {
			switch {
			case <-fetchSig:
				// Check if signal arrived earlier than a second (with error margin)
				if time.Since(start).Nanoseconds() < 950000000 {
					fail <- true
					return
				}
				start = time.Now()
			}
		}
	}(start)

	main()
	
	wg.Wait()
	if <-fail {
		t.Error("There exist a two crawls that were executed less than 1 second apart. Solution is incorrect.")
	}
}
