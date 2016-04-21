package main

import (
	"fmt"
)

func main() {
	xs := []float64{1.1, 2.2, 3.3}
	avg := average(xs)
	fmt.Printf("%v", avg)
}

func average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
		case 0:
			avg = 0.0
		default:
			for _, v := range xs {
				sum += v
			}
			avg = sum / float64(len(xs)) 
	}
	return
}
