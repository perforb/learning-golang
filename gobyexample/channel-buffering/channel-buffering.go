package main

import "fmt"

// 0x14000068120
// buffered
// channel
// 0x14000068120
func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(messages)
}
