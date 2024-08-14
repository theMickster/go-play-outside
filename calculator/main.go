package main

import (
	"main/shape"
)

func main() {
	x := shape.Square{Length: 7}
	shape.GetArea(x)
	shape.GetPerimeter(x)

	y := shape.Rectangle{Length: 4, Width: 2}
	shape.GetArea(y)
	shape.GetPerimeter(y)

	z := shape.Triangle{Base: 2, Height: 4, LeftSide: 4, RightSide: 4}
	shape.GetArea(z)
	shape.GetPerimeter(z)

	a := shape.Circle{Radius: 9}
	shape.GetArea(a)
	shape.GetPerimeter(a)
}
