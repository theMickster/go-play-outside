package shape

import "math"

type Square struct {
	length float64
}

func (s Square) Area() float64 {
	return math.Pow(s.length, 2)
}

func (s Square) Perimeter() float64 {
	return s.length * 4
}

func ConstructSquare(length float64) Square {
	return Square{length}
}
