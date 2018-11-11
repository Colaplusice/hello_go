package main

import (
	"fmt"
	"os"
)

func add(x int,y int)int {
	return x+y
}
func swap(x, y string) (string, string) {
	return y,x
}

func main() {
fmt.Print(len(os.Args))
}


