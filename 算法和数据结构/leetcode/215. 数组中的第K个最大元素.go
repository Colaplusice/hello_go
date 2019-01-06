package main

import (
	"fmt"
	"errors"
	"hello_go/utils"
)

func findKthLargest(nums []int, k int) int {
	nums = QuickSort(nums)
	return nums[len(nums)-k]
	//return nums
}

// 递归快速排序，选一个数字作为基准值，小于基准值的放左边，大于基准值的放右边
func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	flag := nums[len(nums)-1]
	NumLength := len(nums)
	LeftLength := 0
	for i := 0; i < NumLength; i++ {
		if nums[i] < flag {
			utils.ArraySwitch(nums, i, LeftLength)
			LeftLength += 1
		}
	}
	ArraySwitch(nums, len(nums)-1, LeftLength)
	return append(append(QuickSort(nums[:LeftLength]), flag), QuickSort(nums[LeftLength+1:])...)
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

func main() {

	//nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	//nums := []int{7, 10,  8}
	nums := []int{7, 10, 7, 8, 11, 8}

	//nums := []int{6, 5, 4, 3, 2, 1}
	res := findKthLargest(nums, 1)
	fmt.Println(res)
}
