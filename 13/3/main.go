package main

import "fmt"

func main() {
	// операция обмена с помощью XOR( исключающее ИЛИ)
	a := 4
	b := 8
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)
}
