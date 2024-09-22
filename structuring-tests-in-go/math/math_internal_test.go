package math

import (
	"testing"
)

func TestMultipy(t *testing.T) {
	if multipy(2, 3) != 6 {
		t.Fail()
	}
}