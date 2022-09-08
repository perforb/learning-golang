package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// Worker 5 starting
// Worker 3 starting
// Worker 1 starting
// Worker 4 starting
// Worker 2 starting
// Worker 2 done
// Worker 3 done
// Worker 5 done
// Worker 1 done
// Worker 4 done
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// https://go.dev/doc/faq#closures_and_goroutines
		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
