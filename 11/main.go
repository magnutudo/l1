package main

import "fmt"

func main() {
	frstSet := []int{2, 34, 11, 66, 21, 22, 5}
	scSet := []int{2, 1, 5, 6, 1, 11}

	res := Intersec(frstSet, scSet)
	fmt.Println(res)
}
func Intersec(frstSet []int, scSet []int) []int {
	set := make(map[int]bool)
	interction := []int{}
	// отмечаем все элементы 1 множества как true
	for _, v := range frstSet {
		set[v] = true
	}
	// сравниваем элементы второго множества, есть ли они в map`e
	for _, v := range scSet {
		_, ok := set[v]
		if ok == true {
			interction = append(interction, v)
		}
	}
	return interction
}
