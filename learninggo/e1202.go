package main

import (
	"fmt"
)

func Map(f func(string) string, l []string) []string {
	j := make([]string, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}

func main() {
	m := []string{"um", "tres", "quatro"}
	f := func(i string) string {
		return i + " - " + i
	}
	fmt.Printf("%v", (Map(f, m)))
}
