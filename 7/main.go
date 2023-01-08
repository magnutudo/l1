package main

import (
	"fmt"
	"sync"
)

func main() {
	ourMap := make(map[int]string)
	mut := &sync.RWMutex{}
	group := &sync.WaitGroup{}
	for i := 1; i <= 15; i++ {
		// вызываеm Add, чтобы установить количество goroutine, которых необходимо ожидать
		group.Add(1)
		go func(num int) {

			mut.Lock()
			ourMap[num] = "Привет,куратор\n"

			mut.Unlock()
			defer group.Done() // каждая горутина вызывает Done, когда завершается
		}(i)
	}
	fmt.Println(ourMap)
}
