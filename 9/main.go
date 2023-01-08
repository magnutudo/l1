package main

import (
	"fmt"
	"os"
)

func inputToChann(mas []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, elem := range mas {
			out <- elem
		}
		close(out)
	}()
	return out
}
func square(in <-chan int) <-chan int {
	newOut := make(chan int)
	go func() {
		for elem := range in {
			newOut <- elem * 2
		}
		close(newOut)
	}()
	return newOut
}
func main() {
	mas := []int{2, 4, 6, 8, 10}
	inp := inputToChann(mas)
	sq := square(inp)
	fmt.Fprintln(os.Stdout, <-sq)
	fmt.Fprintln(os.Stdout, <-sq)
	fmt.Fprintln(os.Stdout, <-sq)
	fmt.Fprintln(os.Stdout, <-sq)
	fmt.Fprintln(os.Stdout, <-sq)

}
