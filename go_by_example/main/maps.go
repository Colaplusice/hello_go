package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println("len", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

}
