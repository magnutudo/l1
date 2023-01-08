package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	N := 1000
	inp := inputToChan(N)
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	rd := read(inp, ctx)

	for _ = range rd {
		fmt.Println(<-rd)
	}

}
func inputToChan(n int) <-chan int {
	out := make(chan int)
	b := n * 1000
	go func() {
		for i := 0; i < b; i++ {
			out <- i
		}
		close(out)
	}()

	return out

}
func read(in <-chan int, ctx context.Context) <-chan int {
	newOut := make(chan int)
	go func() {
		for {
			select {
			case res := <-in:
				fmt.Println(res)
			case <-ctx.Done():
				os.Exit(0)
			}
		}

		close(newOut)
	}()

	return newOut
}
