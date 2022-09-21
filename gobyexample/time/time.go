package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 2022-09-21 23:04:07.860678 +0900 JST m=+0.000073292
	fmt.Println(now)
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	// 2009-11-17 20:34:58.651387237 +0000 UTC
	fmt.Println(then)
	// 2009
	fmt.Println(then.Year())
	// November
	fmt.Println(then.Month())
	// 17
	fmt.Println(then.Day())
	// 20
	fmt.Println(then.Hour())
	// 34
	fmt.Println(then.Minute())
	// 58
	fmt.Println(then.Second())
	// 651387237
	fmt.Println(then.Nanosecond())
	// UTC
	fmt.Println(then.Location())
	// Tuesday
	fmt.Println(then.Weekday())
	// true
	fmt.Println(then.Before(now))
	// false
	fmt.Println(then.After(now))
	// false
	fmt.Println(then.Equal(now))

	diff := now.Sub(then)
	// 112577h35m33.226039763s
	fmt.Println(diff)
	// 112577.61477035077
	fmt.Println(diff.Hours())
	// 6.754656886221046e+06
	fmt.Println(diff.Minutes())
	// 4.052794131732628e+08
	fmt.Println(diff.Seconds())
	// 405279413173262763
	fmt.Println(diff.Nanoseconds())
	// 2022-09-21 14:11:51.82465 +0000 UTC
	fmt.Println(then.Add(diff))
	// 1997-01-14 02:58:05.478124474 +0000 UTC
	fmt.Println(then.Add(-diff))
}
