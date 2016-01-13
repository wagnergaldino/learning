package main

import "fmt"

func max(args ...int) int {
  max := 0
  for x, v := range args {
    if x == 0 || v > max {
    	max = v
    }
  }
  return max
}
func main() {
  fmt.Println(max(5,1,2,3))
}
