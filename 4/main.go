package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

func main() {
	sgnCh := make(chan os.Signal, 1)
	signal.Notify(sgnCh, os.Interrupt)

	go func() {
		for sig := range sgnCh {
			fmt.Println(sig)
			os.Exit(0)
		}
	}()

	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Error with Scan")
	}
	c := make(chan float64)

	for i := 0; i < N; i++ {
		go func() {
			for {
				fmt.Println(<-c)
			}
		}()
	}

	for {
		c <- rand.Float64()
	}
}
