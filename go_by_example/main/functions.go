package main

import "fmt"

func plus(a int, b int) (int, string) {
	return a, "a"
}

func main() {
	sd, str_sd := plus(2, 3)
	fmt.Println(sd)
	fmt.Println(str_sd)
}
