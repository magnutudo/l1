package main

import "fmt"

func main() {
	a := []int{2, 5, 1, 27, 9}
	index := 2
	deleteSlice := append(a[:index], a[index+1:]...) // Один срез добавляется к другому с помощью оператора ... для расширения исходного среза на отдельные значения ( из книги Идиомы и паттерны проектирования )
	fmt.Println(deleteSlice)
}
