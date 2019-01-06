package main

import (
	"hello_go/utils"
	"fmt"
)

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
	utils.ArraySwitch(nums, len(nums)-1, LeftLength)
	return append(append(QuickSort(nums[:LeftLength]), flag), QuickSort(nums[LeftLength+1:])...)
}
func main() {
	//nums := []int{4, 2, 7, 6, 2, 3, 5, 5, 9}
	nums := []int{7, 10, 7, 8, 11, 8}
	result := QuickSort(nums)
	fmt.Println(result)

}
