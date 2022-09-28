package main

import (
	"flag"
	"fmt"
)

// $ ./command-line-flags -word=opt a1 a2 a3 -numb=7                                             (git)-[master]
// word: opt
// numb: 21
// fork: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]
func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 21, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
