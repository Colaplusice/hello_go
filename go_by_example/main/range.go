package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Print(nums)
	for index, num := range nums {
		fmt.Println(index, num)
	}

	dict := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range dict {
		fmt.Println(key, value)
	}
	for index,st :=range "sdasda"{
		fmt.Println(index,st)
	}
}
