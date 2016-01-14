package main

import "fmt"

func swap(x, y int) (a, b int) {
  a, b = y, x
  return
}

func main() {
  fmt.Println(swap(1, 2))
}
