package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	sum(1,2,34)
}

func sum(nums ...int) {

	for i := range nums {
		fmt.Println(i)
	}

}
