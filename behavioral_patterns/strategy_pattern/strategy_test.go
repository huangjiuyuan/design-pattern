package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	ctx := NewContext(&OperationAdd{})
	sum := ctx.ExecuteStrategy(25, 10)
	if sum != 35 {
		t.Errorf("expected 35, got %d", sum)
	}

	ctx = NewContext(&OperationSubstract{})
	diff := ctx.ExecuteStrategy(25, 10)
	if diff != 15 {
		t.Errorf("expected 15, got %d", diff)
	}

	ctx = NewContext(&OperationMultiply{})
	prod := ctx.ExecuteStrategy(25, 10)
	if prod != 250 {
		t.Errorf("expected 15, got %d", prod)
	}
}
