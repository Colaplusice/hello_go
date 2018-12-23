package main

import "fmt"

func maxArea(height []int) int {
	result := 0

	for index_1, each_1 := range height {
		for i := index_1; i < len(height); i++ {
			temp_value := 0
			if each_1 > height[i] {
				temp_value = (i - index_1) * height[i]
			} else {
				temp_value = (i - index_1) * each_1
			}
			if temp_value > result {
				result = temp_value
			}
		}
	}
	return result
}

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(a)
	fmt.Println(res)
}
