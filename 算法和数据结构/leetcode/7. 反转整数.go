package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {

	a := 0
	for x != 0 {
		a *= 10
		a += x % 10
		x /= 10

		if a > math.MaxInt32 || a < math.MinInt32 {
			return 0

		}
	}
	return a

}

func main() {

	res := reverse(1534236469)
	fmt.Println(res)
}
