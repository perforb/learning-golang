package main

import (
	"fmt"
	"strconv"
)

// 1.234
// 123
// 83
// 123
// 456
// 789
// 135
// strconv.Atoi: parsing "wat": invalid syntax
func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	i2, _ := strconv.ParseInt("000123", 0, 64)
	fmt.Println(i2)
	i3, _ := strconv.ParseInt("000123", 10, 64)
	fmt.Println(i3)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	_, err := strconv.Atoi("wat")
	fmt.Println(err)
}
