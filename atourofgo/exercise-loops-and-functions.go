package main

import (
	"fmt"
	"math"
)

//alterei o retorno da função para incluir um contador
//que registra o numero de iterações no loop
func Sqrt(x float64) (float64, int) {
	z := x
	s := 0.0	
	c := 0
	d := 0.00000001
	for {
		z = z - (z*z - x)/(2*z)
		c += 1
		if math.Abs(s-z) < d {
			break
		}
		s = z
		c += 1
	}
	return s, c
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
	fmt.Println(Sqrt(7))
	fmt.Println(math.Sqrt(7))
	fmt.Println(Sqrt(9))
	fmt.Println(math.Sqrt(9))
	fmt.Println(Sqrt(25))
	fmt.Println(math.Sqrt(25))
}
