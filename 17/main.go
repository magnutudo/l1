package main

import (
	"fmt"
	"strconv"
)

func main() {
	var elems, digit int

	fmt.Print("Number of elements? ")
	fmt.Scan(&elems)

	var array = make([]int, elems)

	fmt.Print("Generated array: ")
	for i := 0; i < elems; i++ {
		array[i] = i
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()

	fmt.Println("Что мы ищем? :")
	fmt.Scan(&digit)
	var assumption, result = binarySearch(&array, digit)
	if result {
		fmt.Println("Найден: " + strconv.Itoa(assumption))
	} else {
		fmt.Println("Такого числа нет")
	}
}
func binarySearch(array *[]int, digit int) (int, bool) {

	var mid, assumption int

	min := 0
	high := len(*array) - 1

	for min <= high {
		mid = (min + high) / 2
		assumption = mid
		if assumption == digit {
			return assumption, true
		}
		if assumption > digit {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return 0, false
}
