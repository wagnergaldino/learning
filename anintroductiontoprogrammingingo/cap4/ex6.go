package main

import "fmt"

func main() {
  fmt.Print("Entre uma distancia em pés: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 0.3048

  fmt.Println(input, " pés = ", output, " metros")
}
