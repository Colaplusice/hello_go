package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 2, height: 10}
	fmt.Print("area", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())

}
