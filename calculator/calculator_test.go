package calculator

import "testing"

func TestCalculateAddValuesSucceeds(t *testing.T) {
	leftInput, rightInput := "7", "25"
	expect := 32.00

	result := Calculate(leftInput, rightInput, "+")

	if expect != result {
		t.Errorf("Failed to add %v and %v. Result %v but expected %v\n", leftInput, rightInput, result, expect)
	}
}
