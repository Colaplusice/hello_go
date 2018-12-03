package main

import "fmt"

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//你可以假设数组中无重复元素。

// 二分法

func searchInsert(nums []int, target int) int {
	middle := 0
	min := 0
	max := len(nums) - 1
	for min <= max {
		middle = (min + max) / 2
		if target == nums[middle] {
			return middle
		} else if target > nums[middle] {
			min = middle + 1
		} else {
			max = middle - 1
		}
	}
return min
}
func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	index := searchInsert(nums, target)
	fmt.Println(index)

}
