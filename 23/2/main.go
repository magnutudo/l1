package main

import "fmt"

func main() {
	sl := []int{2, 6, 1, 12, 88}
	index := 2
	// выполняем сдвиг
	copy(sl[index:], sl[index+1:]) // Встроенная функция копирования копирует элементы из исходного фрагмента в целевой фрагмент
	fmt.Println(sl)
	// усекаем конец
	sl = sl[:len(sl)-1]
	fmt.Println(sl)
}
