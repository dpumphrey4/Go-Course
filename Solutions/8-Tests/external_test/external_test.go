package externaltest

import (
	"testing"
	"tests/calculator"
)

func TestMul(t *testing.T) {
	result := calculator.Mul(2, 3)

	if result != 6 {
		t.Errorf("expected 6, got %d", result)
	}
}
