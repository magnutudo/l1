package main

import (
	"fmt"
)

var counter int

func main() {

	ch := make(chan bool)

	for i := 1; i <= 10; i++ {
		go work(i, ch)
	}

	for i := 1; i <= 10; i++ {
		<-ch
	}

	fmt.Println("конец")
}
func work(number int, ch chan bool) {

	counter++
	fmt.Println("Goroutine", number, "-", counter)
	ch <- true

}
