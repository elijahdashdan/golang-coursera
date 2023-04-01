package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter up to 10 integers:")
	integers := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scanln(&integers[i])
	}

	BubbleSort(integers)
	fmt.Println("Sorted numbers:", integers)
}

func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}
