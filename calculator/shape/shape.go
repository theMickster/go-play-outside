package shape

import (
	"fmt"
	"reflect"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func GetArea(s Shape) {
	fmt.Printf("Area of the %s is: %v \n", reflect.TypeOf(s).Name(), s.Area())
}

func GetPerimeter(s Shape) {
	fmt.Printf("Perimeter of the %s is: %v \n", reflect.TypeOf(s).Name(), s.Perimeter())
}
