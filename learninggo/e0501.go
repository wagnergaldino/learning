package main

import (
	"fmt"
)

func main() {
	sum := 0.0
	xs := []float64{1.1, 2.2, 3.3}
	var avg float64
	switch len(xs) {
		case 0: 
			avg = 0.0
		default: 
			for _, v := range xs {
				sum += v
			}
			avg = sum / float64(len(xs))
	}
	fmt.Printf("%v", avg)
}
