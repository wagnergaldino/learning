package main

import (
	"fmt"
)

func main() {
	xs := []int{1, 2, 3}
	m := min(xs)
	fmt.Printf("%v", m)
}

func min(l []int) (min int) {
	min = l[0]
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	return
}
