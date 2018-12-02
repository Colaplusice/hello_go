package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	lens := len(strs[0])
	for i := 0; i < lens; i++ {
		prefix := string(strs[0][i])
		for index, _ := range strs {
			if i >=len(strs[index]) {
				return strs[index][:i]
			}
			if string(strs[index][i]) != prefix {
				return strs[index][:i]
			}
		}
	}
	return  strs[0]
}

func main() {
	a := []string{"c","c"}
	res := longestCommonPrefix(a)
	fmt.Println(res)

}
