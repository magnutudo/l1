package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	// превращаю строку в массив
	words := strings.Fields(s)
	// меняю порядок первый - последний, уменьшая j
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(Reverse("snow dog sun"))
}
