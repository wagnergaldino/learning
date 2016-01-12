package main

import "fmt"

func main() {
  fmt.Print("Entre uma temperatura em graus Farenheit: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := (input - 32) * 5/9

  fmt.Println(input, " Graus Farenheit = ", output, " Graus Celsius")
}
