package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 10, height: 5}

	// Value receiver method call
	area := rect.Area()
	fmt.Println(area)

	// Pointer receiver method call
	rect.Scale(2)
	fmt.Println(rect)

}
