package main

import (
	"fmt"
	"sync"
)

var (
	counter int
)

func main() {
	var (
		mutex = sync.Mutex{}
		wg    = sync.WaitGroup{}
	)

	const N = 10

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			work(&mutex, j)
		}(i)
	}

	wg.Wait()

	fmt.Println("Done")
}

func work(m *sync.Mutex, j int) {
	m.Lock()
	counter++
	fmt.Printf("Горутина номер %d; counter %d\n", j, counter)
	m.Unlock()
}
