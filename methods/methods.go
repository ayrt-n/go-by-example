package main

import "fmt"

type rect struct {
	width, length int
}

func (r *rect) area() int {
	return r.width * r.length
}

func (r rect) perimeter() int {
	return r.width*2 + r.length*2
}

func main() {
	r := rect{width: 10, length: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perimeter())
}
