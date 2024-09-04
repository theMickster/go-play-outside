package shape

type Triangle struct {
	base      float64
	height    float64
	leftSide  float64
	rightSide float64
}

func (t Triangle) Area() float64 {
	return .5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	return t.rightSide + t.base + t.leftSide
}

func ConstructTriangle(base, height, leftSide, rightSide float64) Triangle {
	return Triangle{base, height, leftSide, rightSide}
}
