package shape

import "math"

type Circle struct {
	radius float64
}

const floatPrecision = 2

func (c Circle) Area() float64 {
	result := math.Pi * math.Pow(c.radius, 2)
	return roundFloatUp(result, floatPrecision)
}

func (c Circle) Perimeter() float64 {
	result := 2 * math.Pi * c.radius
	return roundFloatUp(result, floatPrecision)
}

func roundFloatUp(value float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(value*ratio) / ratio

}

func ConstructCircle(radius float64) Circle {
	return Circle{radius}
}
