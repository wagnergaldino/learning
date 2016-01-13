package main

import "fmt"

func main() {
  fmt.Print("Entre um numero inteiro e veja a sua sequencia de fibonacci: ")
  var input int
  fmt.Scanf("%d", &input)
  
  fmt.Println(0)
  
  f := fibonacci()
	for i := 0; i < input; i++ {
		fmt.Println(f())
	}

}

func fibonacci() func() int {
	a, b := 0,1
	return func() int {
		a, b = b, a+b
		return a
	}
}
