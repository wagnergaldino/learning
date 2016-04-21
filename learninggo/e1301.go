package main

import (
	"fmt"
)

func main() {
	xs := []int{1, 2, 3}
	m := max(xs)
	fmt.Printf("%v", m)
}

func max(l []int) (max int) {
	max = l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return
}
