package main

import "fmt"

func numberOfLines(widths []int, S string) []int {
	lines := 1
	current_length := 0
	for _, str := range S {
		str_length := widths[str-97]
		if current_length+str_length > 100 {
			lines++
			current_length = str_length
		}else {
			current_length+=str_length
		}
	}
	return []int{lines, current_length}
}

func main() {
	width := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	S := "abcdefghijklmnopqrstuvwxyz"
	res:=numberOfLines(width, S)
	fmt.Println(res)

}
