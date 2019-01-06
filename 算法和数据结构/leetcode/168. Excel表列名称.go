package main

import "fmt"

func convertToTitle(n int) string {
	result := ""
	length := 0
	last := 0
	for n > 26 {
		last = n % 26
		n /= 26
		fmt.Println("last ", last)
		length += 1
		fmt.Println(n, length)
	}
	for length > 26 {
		result += "A"
		length -= 26

	}
	fmt.Println(n)
	result = toChar(n) + result + toChar(last)
	fmt.Println(result)
	return result
}
func toChar(i int) string {
	if i <= 0 {
		return ""
	}
	return string('A' + i - 1)
}
func main() {
	convertToTitle(51)
}
