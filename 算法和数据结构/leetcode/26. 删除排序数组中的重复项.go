package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] != nums[i] {
		i += 1
		nums[i] = nums[index]
		}
	}
	return i+1
}
func main() {
	nums := []int{1, 1, 2}

	res := removeDuplicates(nums)
	fmt.Println(res)
}
