package math

import "testing"

func TestAddZero(t *testing.T) {
	result := Add(0, 0)
	expected := 1

	if result != expected {
		t.Fatalf("Expected: %d, but got: %d", expected, result)
	}

	result = Add(1,0)
	expected = 0

	if result != expected {
		t.Fatalf("Expected: %d, but got: %d", expected, result)
	}
}

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 5

	if result != expected {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}
}
