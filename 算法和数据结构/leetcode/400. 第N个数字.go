package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	i := 1
	commands := 1
	for n > i*9*commands {
		n -= i * 9 * commands
		commands *= 10
		i++
	}
	commands -= 1
	divid := n % (i)
	n /= i

	offset := n
	numbers := commands + offset
	str_number := strconv.Itoa(numbers)
	if divid == 0 {
		result, _ := strconv.Atoi(string(str_number[len(str_number)-divid-1]))
		return result
	} else {
		str_number := strconv.Itoa(numbers + 1)
		result, _ := strconv.Atoi(string(str_number[divid-1]))
		return result
	}
}

func main() {
	res := findNthDigit(28)
	fmt.Println(res)
}
