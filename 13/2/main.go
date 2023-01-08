package main

import "fmt"

func main() {
	// способ с умножением и делением
	a := 2
	b := 5
	a = a * b // a = 10
	b = a / b // b = 2
	a = a / b // a = 5
	fmt.Println(a, b)
}
