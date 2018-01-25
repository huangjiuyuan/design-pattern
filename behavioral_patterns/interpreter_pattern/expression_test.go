package interpreter

import (
	"testing"
)

func TestExpression(t *testing.T) {
	robert := NewTerminalExpression("Robert")
	john := NewTerminalExpression("John")
	or := NewOrExpression(robert, john)

	julie := NewTerminalExpression("Julie")
	married := NewTerminalExpression("Married")
	and := NewAndExpression(julie, married)

	if !or.Interpret("Robert") {
		t.Errorf("expected true, got a false")
	}
	if !or.Interpret("John") {
		t.Errorf("expected true, got a false")
	}
	if !and.Interpret("Married Julie") {
		t.Errorf("expected true, got a false")
	}
}
