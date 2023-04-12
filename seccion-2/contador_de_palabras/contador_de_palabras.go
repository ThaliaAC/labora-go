package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	v := strings.Fields(s)
	words := make(map[string]int)
	for _, value := range v {
		count, ok := words[value]
		if !ok {
			count = 0
		}
		words[value] = count + 1
	}
	return words
}

func main() {
	s := "I ate a donut. Then I ate another donut."
	v := WordCount(s)
	fmt.Println(v)
}
