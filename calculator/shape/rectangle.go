package shape

type Rectangle struct {
	length float64
	width  float64
}

func (s Rectangle) Area() float64 {
	return s.length * s.width
}

func (s Rectangle) Perimeter() float64 {
	return 2 * (s.length + s.width)
}

func ConstructRectangle(length, width float64) Rectangle {
	return Rectangle{length, width}
}
