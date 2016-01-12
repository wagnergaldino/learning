package main

import "fmt"

func main() {
  i := 1
  for i <= 10 {
  	if i % 2 == 0 {
		  fmt.Println(i, " é par")
		} else {
		  fmt.Println(i, " é impar")
		}
    i = i + 1
  }
}
