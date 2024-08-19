package main

import (
	"main/shape"
)

func main() {
	x := shape.ConstructSquare(7)
	shape.GetArea(x)
	shape.GetPerimeter(x)

	y := shape.ConstructRectangle(4, 2)
	shape.GetArea(y)
	shape.GetPerimeter(y)

	z := shape.ConstructTriangle(2, 4, 4, 4)
	shape.GetArea(z)
	shape.GetPerimeter(z)

	a := shape.ConstructCircle(9)
	shape.GetArea(a)
	shape.GetPerimeter(a)
}
