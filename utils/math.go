package utils

import "errors"

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func ArraySwitch(nums []int, index1 int, index2 int) error {
	NumLength := len(nums)
	if NumLength <= index1 || NumLength <= index2 {
		return errors.New("array index out of bound")
	}
	tmp := nums[index1]
	nums[index1] = nums[index2]
	nums[index2] = tmp
	return nil
}
