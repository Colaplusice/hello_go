package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n < 0 {
		return ""
	}
	str := "1"
	newStr := ""
	previous := "1"
	for i := 1; i < n; i++ {
		all := 0
		for _, value := range str {
			if previous == string(value) {
				all += 1
			} else {
				newStr += strconv.Itoa(all) + previous
				all = 1
				previous = string(value)
			}
		}
		newStr += strconv.Itoa(all) + previous
		str = newStr
		newStr = ""
		previous = string(str[0])
	}
	return str
}

func main() {
	res := countAndSay(4)
	fmt.Println(res)

}
