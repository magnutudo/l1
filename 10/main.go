package main

import (
	"fmt"
	"math"
)

func main() {
	tempArr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	group(tempArr)
}
func group(arr []float64) {
	group := map[float64][]float64{}
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			var a float64
			a = math.Ceil(arr[i]/10) * 10
			group[a] = append(group[a], arr[i])
		}
		if arr[i] > 0 {
			a := math.Floor(arr[i]/10) * 10
			group[a] = append(group[a], arr[i])
		}

	}
	fmt.Println(group)
}
