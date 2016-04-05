package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _, word := range words {
		count[word]++
	}
	return count
}

func main() {
	retorno := WordCount("the book is on the table")
	fmt.Printf("%v", retorno)
}
