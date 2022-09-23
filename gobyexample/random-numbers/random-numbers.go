package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"time"
)

func main() {
	byMath()
	byCrypto()
}

func byMath() {
	// 81,87
	// 0.6645600532184904
	// 7.188570935934901,7.123187485356328
	// 出力は固定
	fmt.Print(mrand.Intn(100), ",")
	fmt.Print(mrand.Intn(100))
	fmt.Println()
	fmt.Println(mrand.Float64())
	fmt.Print((mrand.Float64()*5)+5, ",")
	fmt.Print((mrand.Float64() * 5) + 5)
	fmt.Println()

	// random numbers
	s1 := mrand.NewSource(time.Now().UnixNano())
	r1 := mrand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 5,87
	s2 := mrand.NewSource(42)
	r2 := mrand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	// 5,87
	s3 := mrand.NewSource(42)
	r3 := mrand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
	fmt.Println()
}

// https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
// https://qiita.com/crifff/items/b116e6235fedcd18e0de
func byCrypto() {
	// 65
	n, _ := crand.Int(crand.Reader, big.NewInt(100))
	fmt.Println(n)
}
