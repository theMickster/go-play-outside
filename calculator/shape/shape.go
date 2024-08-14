package shape

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	Length float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base      float64
	Height    float64
	LeftSide  float64
	RightSide float64
}

func (s Square) Area() float64 {
	return math.Pow(s.Length, 2)
}

func (s Square) Perimeter() float64 {
	return s.Length * 4
}

func (s Rectangle) Area() float64 {
	return s.Length * s.Width
}

func (s Rectangle) Perimeter() float64 {
	return 2 * (s.Length + s.Width)
}

func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.RightSide + t.Base + t.LeftSide
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func GetArea(s Shape) {
	fmt.Printf("Area of the %s is: %v \n", reflect.TypeOf(s).Name(), s.Area())
}

func GetPerimeter(s Shape) {
	fmt.Printf("Perimeter of the %s is: %v \n", reflect.TypeOf(s).Name(), s.Perimeter())
}
