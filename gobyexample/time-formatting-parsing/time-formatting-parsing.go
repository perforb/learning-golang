package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// 2022-09-22T22:55:36+09:00
	fmt.Println(t.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2021-11-01T22:08:41+00:00")
	// 2021-11-01 22:08:41 +0000 +0000
	fmt.Println(t1)
	// 10:58PM
	fmt.Println(t.Format("3:04PM"))
	// Thu Sep 22 22:58:48 2022
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	// 2022-09-22T23:00:03.288584+09:00
	fmt.Println(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	// 0000-01-01 20:41:00 +0000 UTC
	fmt.Println(t2)

	// 2022-09-22T23:04:14-00:00
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	// parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
	fmt.Println(e)
}
