package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	num_len := len(nums)
	i := 0
	for index := 0; index < num_len-i; index += 1 {
		fmt.Println("this is index", index)
		if val == nums[index] {
			fmt.Println("this val", val)
			if index+1 < num_len {
				nums[index] = nums[index+1]
			}
			i += 1
			index -= 1
		}
	}
	fmt.Println("this is i", i)
	nums = nums[:num_len-i]
	fmt.Println(nums)
	return num_len - i
}

func main() {
	nums := []int{3, 2, 2, 3}
	res := removeElement(nums, 3)
	fmt.Println(res)
	fmt.Println(nums)
}
