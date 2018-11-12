package main

import "fmt"

func addDigits(num int) int {
	a := 0
	for num >= 10 {
		for num > 0 {
			a += num % 10
			num /= 10
		}
		num = a
		a = 0
	}
	return num
}
func main() {
	res:=addDigits(10)
	fmt.Println(res)

}
