package main

import (
	"fmt"
	"time"
)

// request 1 2022-09-09 17:35:04.029889 +0900 JST m=+0.201120334
// request 2 2022-09-09 17:35:04.232555 +0900 JST m=+0.403784251
// request 3 2022-09-09 17:35:04.429591 +0900 JST m=+0.600819043
// request 4 2022-09-09 17:35:04.629885 +0900 JST m=+0.801111376
// request 5 2022-09-09 17:35:04.829305 +0900 JST m=+1.000529251
//
// request 1 2022-09-09 17:35:04.829489 +0900 JST m=+1.000712626
// request 2 2022-09-09 17:35:04.8295 +0900 JST m=+1.000724209
// request 3 2022-09-09 17:35:04.829508 +0900 JST m=+1.000731876
// request 4 2022-09-09 17:35:05.029894 +0900 JST m=+1.201116043
// request 5 2022-09-09 17:35:05.229657 +0900 JST m=+1.400877334
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println()

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
