package main

import (
	"fmt"
)

func maxArea(height []int) int {
	begin := 0
	end := len(height) - 1
	result := mini(height[begin], height[end]) * end
	for begin < end {
		tempResult := 0
		if height[begin] < height[end] {
			begin += 1
			tempResult = mini(height[begin], height[end]) * (end - begin)
		} else {
			end -= 1
			tempResult = mini(height[begin], height[end]) * (end - begin)
		}
		if tempResult > result {
			result = tempResult
		}
	}
	return result
}
func mini(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	nums := [] int{2, 3, 4, 5, 18, 17, 6}
	res := maxArea(nums)
	fmt.Print(res)
}
