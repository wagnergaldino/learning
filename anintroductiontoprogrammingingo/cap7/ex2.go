package main

import "fmt"

func half(z int) (j int, k bool) {
  j = z/2
  k = z%2 == 0
  return
}

func main() {
  x, y := half(2)
  fmt.Println(x, y)
}
