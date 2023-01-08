package main

import "fmt"

func main() {
	// Арифметический способ обмена значениями
	a := 2
	b := 5
	a = a + b // a = 7
	b = a - b // b = 2
	a = a - b // 7 - 2 = 5
	fmt.Println(a, b)
}
