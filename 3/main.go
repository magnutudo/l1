package main

import (
	"fmt"
	"math"
	"os"
)

var mas = []float64{2, 4, 6, 8}

func main() {
	chanForSum := make(chan float64)
	go square(chanForSum)
	fmt.Fprintln(os.Stdout, <-chanForSum)
}

func square(ch chan float64) {
	var sum float64

	for _, v := range mas {
		sum += math.Pow(v, 2)
	}

	ch <- sum
}
