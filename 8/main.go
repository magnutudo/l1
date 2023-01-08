package main

import "fmt"

func changeBit(num *int64, index uint8, bit uint8) {
	fmt.Printf("%b\n", *num)
	if bit == 1 {
		*num = *num | (1 << index)
	} else if bit == 0 {
		*num = *num & (1<<63 - 1 - (1 << index))
	} else {
		fmt.Println("Неправильный бит")
		return
	}
	fmt.Printf("%b\n", *num)
}

func main() {
	var num int64 = 1<<15 - 1
	changeBit(&num, 7, 0)
}
