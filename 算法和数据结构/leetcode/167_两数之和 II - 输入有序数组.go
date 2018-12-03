package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	end := len(numbers) - 1

	for begin := 0; begin < end; {
		if numbers[begin]+numbers[end] == target {
			return []int{begin + 1, end + 1}
		}
		if numbers[begin]+numbers[end] < target {
			begin += 1
		} else {
			end -= 1
		}
	}
	return []int{}
}

func main() {
	a := []int{2, 7, 11, 15}
	b := 9
	res := twoSum(a, b)
	fmt.Println(res)
}
