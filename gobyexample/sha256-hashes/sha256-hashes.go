package main

import (
	"crypto/sha256"
	"fmt"
)

// sha256 this string
// 1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb
func main() {
	s := "sha256 this string"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
