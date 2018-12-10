package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func test(person *Person) {
	print(person)
	fmt.Println(person.age)
}
func main() {
	fmt.Print(Person{"bob", 20})
	s := Person{"lua", 23}
	s.age = 30
	fmt.Println(s)
	test(&s)
}
