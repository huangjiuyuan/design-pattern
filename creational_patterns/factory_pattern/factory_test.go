package shape

import (
	"reflect"
	"testing"
)

func TestGetShape(t *testing.T) {
	f := new(Factory)

	defaultRect := &rect{
		shapeType: "rect",
		width:     20,
		height:    15,
	}
	rect := f.GetShape("rect")
	rect.Draw()
	if !reflect.DeepEqual(defaultRect, rect) {
		t.Fatalf("expected rect: %#v, got a different: %#v", defaultRect, rect)
	}

	defaultSquare := &square{
		shapeType: "square",
		width:     15,
	}
	square := f.GetShape("square")
	square.Draw()
	if !reflect.DeepEqual(defaultSquare, square) {
		t.Fatalf("expected square: %#v, got a different: %#v", defaultSquare, square)
	}

	defaultCircle := &circle{
		shapeType: "circle",
		radius:    10,
	}
	circle := f.GetShape("circle")
	circle.Draw()
	if !reflect.DeepEqual(defaultCircle, circle) {
		t.Fatalf("expected circle: %#v, got a different: %#v", defaultCircle, circle)
	}

	defaultEmpty := &empty{}
	empty := f.GetShape("empty")
	empty.Draw()
	if !reflect.DeepEqual(defaultEmpty, empty) {
		t.Fatalf("expected empty: %#v, got a different: %#v", defaultEmpty, empty)
	}
}
