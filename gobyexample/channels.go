package main

import "fmt"

// 0x1400005c060
// ping
func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	fmt.Println(messages)
	msg := <-messages
	fmt.Println(msg)
}
