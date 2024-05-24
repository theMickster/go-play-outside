package calculator

import (
	"math"
	"testing"
)

func TestCalculateAddValuesSucceeds(t *testing.T) {
	leftInput, rightInput := "7", "25"
	expect := 32.00

	result := Calculate(leftInput, rightInput, "+")

	if expect != result {
		t.Errorf("Failed to add %v and %v. Result %v but expected %v\n", leftInput, rightInput, result, expect)
	}
}

func TestCalculateSubtractValuesSucceeds01(t *testing.T) {
	leftInput, rightInput := "25", "7"
	expect := 18.00

	result := Calculate(leftInput, rightInput, "-")
	if expect != result {
		t.Errorf("Failed to subtract %v and %v. Result %v but expected %v\n", leftInput, rightInput, result, expect)
	}
}

func TestCalculateSubtractValuesSucceeds02(t *testing.T) {
	leftInput, rightInput := "7", "25"
	expect := -18.00

	result := Calculate(leftInput, rightInput, "-")
	if expect != result {
		t.Errorf("Failed to subtract %v and %v. Result %v but expected %v\n", leftInput, rightInput, result, expect)
	}
}

func TestCalculateMultiplyValuesSucceeds(t *testing.T) {
	leftInput, rightInput := "7", "5"
	expect := 35.00

	result := Calculate(leftInput, rightInput, "*")

	if expect != result {
		t.Errorf("Failed to multiply %v and %v. Result %v but expected %v\n", leftInput, rightInput, result, expect)
	}
}

func TestCalculateDivideValuesHappyPathSucceed(t *testing.T) {
	leftInput, rightInput := "28", "7"
	expect := 4.00

	result := Calculate(leftInput, rightInput, "/")

	if expect != result {
		t.Errorf("Failed to divide %v and %v. Result %v but expected %v\n", leftInput, rightInput, result, expect)
	}
}

func TestCalculateDivideByZeroSucceed(t *testing.T) {
	leftInput, rightInput := "28", "0"
	expect := math.Inf(1)

	result := Calculate(leftInput, rightInput, "/")

	if expect != result {
		t.Errorf("Failed to divide %v and %v. Result %v but expected %v\n", leftInput, rightInput, result, expect)
	}
}
