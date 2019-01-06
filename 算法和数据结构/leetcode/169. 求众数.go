package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	many,values:=1,nums[0]
	for _, value := range nums[1:] {
		if value==values{
			many++
		}else {
			many--
			if many<=0{
				many=1
				values=value
			}
		}
	}
	return values
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := majorityElement(nums)
	fmt.Println(res)
}
