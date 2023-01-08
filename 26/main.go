package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "AbcdE"
	lowstr := strings.ToLower(str1) // приведение к 1 регистру
	fmt.Println(uniqueString(lowstr))
}

// создание 2 циклов - один анализирует i элемент, другой - i+1, в случае совпадения - остановка цикла
func uniqueString(str string) (string, bool) {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return str, false
			}
		}
	}
	return str, true
}
