package facade

import (
	"testing"
)

func TestShapeMaker(t *testing.T) {
	sm := NewShapeMaker()
	sm.DrawRect()
	sm.DrawCircle()
}
