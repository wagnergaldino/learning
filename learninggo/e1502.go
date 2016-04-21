package main

import (
	"fmt"
)

func main() {
	p1 := plusX(x)
	p2 := plusX(y)
	fmt.Printf("%v\n",p1(2))
	fmt.Printf("%v\n",p2(2))
}

var x = 2
var y = 5

func plusX(x int) func(int) int {
	return func(y int) int { return x + y }
}
