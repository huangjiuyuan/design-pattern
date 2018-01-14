package bridge

import (
	"testing"
)

func TestDraw(t *testing.T) {
	rc := NewCircle(100, 100, 100, &RedCircle{})
	gc := NewCircle(50, 50, 50, &GreenCircle{})
	rc.Draw()
	gc.Draw()
}
