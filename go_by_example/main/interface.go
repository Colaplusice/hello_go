package main

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, hight float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.hight
}

func (r rect) perim() float64 {
	return 2*r.width
}
