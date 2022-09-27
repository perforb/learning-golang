package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// $ cat /tmp/lines | go run line-filters.go
// HELLO
// FILTER
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
