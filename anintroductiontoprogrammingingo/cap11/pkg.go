package main

import "fmt"
import "github.com/wagnergaldino/learning/anintroductiontoprogrammingingo/cap11/math"

func main() {
  xs := []float64{1,2,3,4}
  avg := math.Average(xs)
  fmt.Println(avg)
  fmt.Println(math.Min(xs))
  fmt.Println(math.Max(xs))
}
