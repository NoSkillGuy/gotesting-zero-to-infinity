package math_test

import (
	"testing"

	customMath "github.com/NoSkillGuy/gotesting-zero-to-infinity/structuring-tests-in-go/math"
)

func TestMultipy(t *testing.T) {
	if customMath.Multiply(2, 3) != 6 {
		t.Fail()
	}
}

func TestSquare(t *testing.T) {
	tests := []struct {
		name string
		input int
		expected int
	}{
		{ name: "Square of 2", input: 2, expected: 4 },
		{ name: "Square of -2", input: -2, expected: 4 },
		{ name: "Square of 0", input: 0, expected: 0 },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if customMath.Square(tt.input) != tt.expected {
				t.Errorf("expected %d but got %d", tt.expected, customMath.Square(tt.input))
			}
		})
	}
}