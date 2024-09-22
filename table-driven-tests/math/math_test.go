package math

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name 	 string
		a, b int
		expected int
	}{
		{"Positive numbers",1, 2,3},
		{"Negative numbers",-1, -1,-2},
		{"Mixed numbers",1, -1,0},
		{"Zero",0, 0,0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name 	 string
		a, b int
		expected int
		errExpected bool
	}{
		{"Positive numbers",10, 2,5,false},
		{"Negative numbers",-10, -2,5,false},
		{"Mixed numbers",10, -2,-5,false},
		{"Zero",0, 10,0,true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if tt.errExpected && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tt.errExpected && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}