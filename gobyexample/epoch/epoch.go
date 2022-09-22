package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 2022-09-22 18:04:23.146234 +0900 JST m=+0.000054168
	fmt.Println(now)
	// 1663837589
	fmt.Println(now.Unix())
	// 1663837589974
	fmt.Println(now.UnixMilli())
	// 1663837589974551000
	fmt.Println(now.UnixNano())
	// 2022-09-22 18:06:29 +0900 JST
	fmt.Println(time.Unix(now.Unix(), 0))
	// 2022-09-22 18:06:29.974551 +0900 JST
	fmt.Println(time.Unix(0, now.UnixNano()))
	// 2075-06-14 03:12:58.974551 +0900 JST
	fmt.Println(time.Unix(now.Unix(), now.UnixNano()))
}
