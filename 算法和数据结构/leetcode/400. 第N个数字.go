package main

import "fmt"

func findNthDigit(n int) int {
	i := 1
	commands := 1
	for n > i*9*commands {
		n -= i * 9 * commands
		commands *= 10
		i++
	}
	n /= i
	offset := n
	divid := n % (i)
	numbers := commands + offset - 1
	return numbers % (1 * (divid) * 10)
}

func main() {
	res := findNthDigit(3)
	fmt.Println(res)
}
