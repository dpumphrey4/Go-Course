package calculator

import "testing"

func TestA(t *testing.T) {
	result := Add(2, 3)

	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}
