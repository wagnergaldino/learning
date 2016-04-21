package main

import (
	"fmt"
)

func main() {
	x, s := 7, 2
	a, b := order(x, s)
	fmt.Printf("%v", a, "%v", b)
}

func order(a, b int) (int, int) {
	if a > b {
		return b,a
	}
	return a,b
}
