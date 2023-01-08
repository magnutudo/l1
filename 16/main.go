package main

import "fmt"

func main() {
	nums := []int{2, 321, 32, 1, -2, 5}
	nums = QuickSort(nums, 0, len(nums)-1)
	fmt.Print(nums)
}
func QuickSort(nums []int, low int, high int) []int {
	if low < high {
		pi := Partit(nums, low, high)
		nums = QuickSort(nums, low, pi-1)
		nums = QuickSort(nums, pi+1, high)
	}
	return nums
}

func Partit(nums []int, low int, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j <= high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]

		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
