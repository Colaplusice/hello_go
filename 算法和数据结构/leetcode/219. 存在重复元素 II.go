package main

import "fmt"

//给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
// 使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

func containsNearbyDuplicate(nums []int, k int) bool {
	result := map[int]int{}
	for index, num := range nums {
		if value, exist := result[num]; exist {
			if index-value <= k {
				return true
			}
		}
			result[num] = index
	}
	return false
}

func main() {
	nums := []int{1, 0, 1, 1}
	k := 1
	result := containsNearbyDuplicate(nums, k)
	fmt.Print(result)

}
