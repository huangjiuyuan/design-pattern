package builder

import (
	"testing"
)

func TestMeal(t *testing.T) {
	mb := new(MealBuilder)
	vm := mb.PrepareVegMeal()
	vm.ShowItems()
	vmc := vm.GetCost()
	if vmc != 40 {
		t.Error("Expected 40, got", vmc)
	}

	nm := mb.PrepareNonVegMeal()
	nm.ShowItems()
	nmc := nm.GetCost()
	if nmc != 50 {
		t.Error("Expected 50, got", nmc)
	}
}
