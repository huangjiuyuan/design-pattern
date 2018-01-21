package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	rect := &rect{"Rect", 10, 20}
	redRect := NewShapeDecorator(rect)
	redRect.Draw()

	circle := &circle{"circle", 15}
	redCircle := NewShapeDecorator(circle)
	redCircle.Draw()
}
