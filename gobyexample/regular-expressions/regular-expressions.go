package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	// true
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	// true
	fmt.Println(r.MatchString("peach"))
	// peach
	fmt.Println(r.FindString("peach punch"))
	// idx: [0 5]
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	// [peach ea]
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// [0 5 1 3]
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// [peach punch pinch]
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	// [peach punch]
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// true
	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	// regexp: p([a-z]+)ch
	fmt.Println("regexp:", r)
	// a <fruit>
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	// a PEACH
	fmt.Println(string(out))
}
