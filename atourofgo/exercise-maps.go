package main

import (
	"strings"
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	palavras := strings.Fields(s)
	mapPalavras := make(map[string]int)
	for _, palavra := range palavras {
		contador, encontrado := mapPalavras[palavra]
		if encontrado {
			mapPalavras[palavra] = contador + 1
		}else {
			mapPalavras[palavra] = 1
		}
	}
	return mapPalavras
}

func main() {
	wc.Test(WordCount)
}
