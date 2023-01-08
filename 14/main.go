package main

import (
	"fmt"
)

func main() {
	var intCh chan int
	var v interface{} = intCh
	someType(v)
}

func someType(typ interface{}) {
	switch x := typ.(type) {
	default:
		fmt.Printf("unexpected type %s\n", x)
	case bool:
		fmt.Printf("boolean %t\n", x)
	case chan int:
		fmt.Println("channel")
	case int:
		fmt.Printf("integer %d\n", x)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *x)
	case *int:
		fmt.Printf("pointer to integer %d\n", x)
	}
}
