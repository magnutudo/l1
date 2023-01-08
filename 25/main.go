package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("сейчас будем спать 10 сек, а потом проснемся")
	myOwnSleep(10 * time.Second)
	fmt.Println("Oh shit here we go again")
}
func myOwnSleep(d time.Duration) {
	<-time.After(d)
}
