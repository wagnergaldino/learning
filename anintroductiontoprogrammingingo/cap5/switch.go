package main

import "fmt"

func main() {
  i := 0
  for i <= 10 {  
  	switch i {
			case 0: fmt.Println(i, " Zero")
			case 1: fmt.Println(i, " Um")
			case 2: fmt.Println(i, " Dois")
			case 3: fmt.Println(i, " Tres")
			case 4: fmt.Println(i, " Quatro")
			case 5: fmt.Println(i, " Cinco")
			case 6: fmt.Println(i, " Seis")
			case 7: fmt.Println(i, " Sete")
			case 8: fmt.Println(i, " Oito")
			case 9: fmt.Println(i, " Nove")
			default: fmt.Println("Numero Desconhecido")
		}
    i = i + 1
  }
}
