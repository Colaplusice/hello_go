package main

import "fmt"

func containsDuplicate(nums []int) bool {
	value_dict := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := value_dict[nums[i]]; !ok {
			value_dict[nums[i]] = 1
		} else {
			return true
		}

	}
	return false
}

func main() {
	nums := []int{1, 4, 3, 4, 5}
	res := containsDuplicate(nums)
	fmt.Print(res)

}
