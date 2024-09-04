package shape

import (
	"testing"
)

func TestConstructSquare(t *testing.T) {
	length := float64(7)
	expect := Square{length: length}

	result := ConstructSquare(length)

	if result.length != expect.length {
		t.Errorf("Expected %v when constructing a square. Got %v \n", expect, result)
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{name: "Length == 1", input: 1, expected: 1},
		{name: "Length == 2", input: 2, expected: 4},
		{name: "Length == 3", input: 3, expected: 9},
		{name: "Length == 8", input: 8, expected: 64},
		{name: "Length == 12", input: 12, expected: 144},
	}

	for _, tc := range tests {
		result := ConstructSquare(tc.input).Area()
		if result != tc.expected {
			t.Errorf("%s :: expected %v when calculating area of a square, got %v", tc.name, tc.expected, result)
		}
	}
}

func TestPerimeter(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{name: "Length == 1", input: 1, expected: 4},
		{name: "Length == 2", input: 2, expected: 8},
		{name: "Length == 5", input: 5, expected: 20},
		{name: "Length == 18", input: 18, expected: 72},
		{name: "Length == 24", input: 24, expected: 96},
	}

	for _, tc := range tests {
		result := ConstructSquare(tc.input).Perimeter()
		if result != tc.expected {
			t.Errorf("%s :: expected %v when calculating perimeter of a square, got %v", tc.name, tc.expected, result)
		}
	}
}
