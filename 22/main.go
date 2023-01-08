package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	var y big.Int
	x.SetString("2112213123213123", 10)
	y.SetString("214124121112333", 10)
	var z big.Int
	z.Mul(&x, &y)
	fmt.Print("Умножение: ")
	fmt.Println(z.String())
	z.Add(&x, &y)
	fmt.Print("Складывание: ")
	fmt.Println(z.String())
	fmt.Print("Вычитание: ")
	z.Sub(&x, &y)
	fmt.Println(z.String())
}
