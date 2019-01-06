package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	value_1 := getMaxValue(nums[:len(nums)-1], true)
	value_2 := getMaxValue(nums, false)
	return max(value_1, value_2)
}

func getMaxValue(nums []int, first bool) int {
	result := make([]int, len(nums))
	if !first {
		result[0] = 0
		result[1] = nums[1]
	} else {
		result[0] = nums[0]
		result[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		//更新result最大值
		if nums[i]+result[i-2] > result[i-1] {
			result[i] = nums[i] + result[i-2]
		} else {
			result[i] = result[i-1]
		}
	}
	return result[len(nums)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, 2}
	res := rob(nums)
	fmt.Println(res)

}
