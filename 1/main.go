package main

import "fmt"

type Human struct {
	Height float64
	Weight float64
}

func (h Human) printWeight() {
	fmt.Printf("Твой вес: %v\n", h.Weight)
}
func (h Human) printHeight() {
	fmt.Printf("Твой рост: %v\n", h.Height)
}

type Action struct {
	Human
}

func main() {
	act := Action{Human{179, 62}}
	act.printWeight()
	act.printHeight()
}
