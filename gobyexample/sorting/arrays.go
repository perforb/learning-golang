package main

import (
	"fmt"
	"sort"
)

type T [2]int

func (t *T) Len() int           { return len(t) }
func (t *T) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t *T) Less(i, j int) bool { return t[i] < t[j] }

// https://stackoverflow.com/questions/20998394/go-sorts-slice-correctly-but-doesnt-sort-array
func main() {
	var x = T([2]int{1, 0})
	fmt.Println(x)
	sort.Sort(&x)
	fmt.Println(x)
}
