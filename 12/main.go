package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := make([]string, 0)
	for _, elem := range arr {
		if existence(elem, res) {
			// если строки такой нет, то вставляем в множество
			res = append(res, elem)
		}
	}
	fmt.Println(res)
}

// проверка на наличие строки в слайсе
func existence(str string, masStr []string) bool {
	for _, word := range masStr {
		if str == word {
			return false
		}
	}
	return true
}
