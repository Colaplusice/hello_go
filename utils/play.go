package main

import "errors"

func main() {
	a := []int{1, 2, 3}
	value, ok := Id(a, 5)
}

func Id(nums []int, index int) (f int, error) {
	if index > len(nums) {
		return 23, errors.New("math length < nums")
	}
	return 1,nil
}
