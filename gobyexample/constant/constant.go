package main

import (
	"fmt"
	"math"
)

const constant string = "constant"

func main() {
	fmt.Println(constant)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
