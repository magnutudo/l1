package main

import "fmt"

func main() {
	a := "главрыбаdawdawdawda↓♠ dwadawd dd →Р5V"
	ReverseString := Reverse(a)
	fmt.Println(ReverseString)
}
func Reverse(s string) string {
	// переводим в руны
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i] // парный обмен
	}
	return string(b)
}
