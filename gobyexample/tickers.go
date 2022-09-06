package main

import (
	"fmt"
	"time"
)

// Tick at 2022-09-06 15:27:14.920373 +0900 JST m=+0.501143626
// Tick at 2022-09-06 15:27:15.420378 +0900 JST m=+1.001150126
// Tick at 2022-09-06 15:27:15.920375 +0900 JST m=+1.501148043
// Ticker stopped
func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
