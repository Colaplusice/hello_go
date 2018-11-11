package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p person
	p.name = "sd"
	p.age = 23
	fmt.Print(p)

}
