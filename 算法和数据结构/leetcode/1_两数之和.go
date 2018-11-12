package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
		if _, ok := m[target-value]; ok {
			return []int{index, m[target-value]}
		}else{
					m[value] = index

		}
	}
	return []int{}
}

func main() {

	a := []int{2, 7, 11, 15}
	b := 9
	res:=twoSum(a, b)
	fmt.Println(res)
}
