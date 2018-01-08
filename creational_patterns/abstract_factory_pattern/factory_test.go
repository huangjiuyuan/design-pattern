package factory

import (
	"reflect"
	"testing"
)

func TestGetFactory(t *testing.T) {
	fp := new(FactoryProducer)
	sf := fp.GetFactory("shape")
	cf := fp.GetFactory("color")

	defaultRect := &rect{
		shapeType: "rect",
		width:     20,
		height:    15,
	}
	rect := sf.GetShape("rect")
	rect.Draw()
	if !reflect.DeepEqual(defaultRect, rect) {
		t.Fatalf("expected rect: %#v, got a different: %#v", defaultRect, rect)
	}

	defaultSquare := &square{
		shapeType: "square",
		width:     15,
	}
	square := sf.GetShape("square")
	square.Draw()
	if !reflect.DeepEqual(defaultSquare, square) {
		t.Fatalf("expected square: %#v, got a different: %#v", defaultSquare, square)
	}

	defaultCircle := &circle{
		shapeType: "circle",
		radius:    10,
	}
	circle := sf.GetShape("circle")
	circle.Draw()
	if !reflect.DeepEqual(defaultCircle, circle) {
		t.Fatalf("expected circle: %#v, got a different: %#v", defaultCircle, circle)
	}

	defaultEmptyShape := &emptyShape{}
	emptyShape := sf.GetShape("empty")
	emptyShape.Draw()
	if !reflect.DeepEqual(defaultEmptyShape, emptyShape) {
		t.Fatalf("expected empty: %#v, got a different: %#v", defaultEmptyShape, emptyShape)
	}

	defaultEmptyRed := sf.GetColor("red")
	emptyRed := sf.GetColor("red")
	emptyRed.Fill()
	if !reflect.DeepEqual(defaultEmptyRed, emptyRed) {
		t.Fatalf("expected empty: %#v, got a different: %#v", defaultEmptyRed, emptyRed)
	}

	defaultEmptyGreen := sf.GetColor("green")
	emptyGreen := sf.GetColor("green")
	emptyGreen.Fill()
	if !reflect.DeepEqual(defaultEmptyGreen, emptyGreen) {
		t.Fatalf("expected empty: %#v, got a different: %#v", defaultEmptyGreen, emptyGreen)
	}

	defaultEmptyBlue := sf.GetColor("blue")
	emptyBlue := sf.GetColor("blue")
	emptyBlue.Fill()
	if !reflect.DeepEqual(defaultEmptyBlue, emptyBlue) {
		t.Fatalf("expected empty: %#v, got a different: %#v", defaultEmptyBlue, emptyBlue)
	}

	defaultRed := &red{
		color: "red",
		red:   255,
		green: 0,
		blue:  0,
	}
	red := cf.GetColor("red")
	red.Fill()
	if !reflect.DeepEqual(defaultRed, red) {
		t.Fatalf("expected red: %#v, got a different: %#v", defaultRed, red)
	}

	defaultGreen := &green{
		color: "green",
		red:   0,
		green: 255,
		blue:  0,
	}
	green := cf.GetColor("green")
	green.Fill()
	if !reflect.DeepEqual(defaultGreen, green) {
		t.Fatalf("expected green: %#v, got a different: %#v", defaultGreen, green)
	}

	defaultBlue := &blue{
		color: "blue",
		red:   0,
		green: 0,
		blue:  255,
	}
	blue := cf.GetColor("blue")
	blue.Fill()
	if !reflect.DeepEqual(defaultBlue, blue) {
		t.Fatalf("expected blue: %#v, got a different: %#v", defaultBlue, blue)
	}

	defaultEmptyColor := &emptyColor{}
	emptyColor := sf.GetColor("empty")
	emptyColor.Fill()
	if !reflect.DeepEqual(defaultEmptyColor, emptyColor) {
		t.Fatalf("expected empty: %#v, got a different: %#v", defaultEmptyColor, emptyColor)
	}

	defaultEmptyRect := cf.GetShape("rect")
	emptyRect := cf.GetShape("rect")
	emptyRect.Draw()
	if !reflect.DeepEqual(defaultEmptyRect, emptyRect) {
		t.Fatalf("expected empty: %#v, got a different: %#v", defaultEmptyRect, emptyRect)
	}

	defaultEmptySquare := cf.GetShape("square")
	emptySquare := cf.GetShape("square")
	emptySquare.Draw()
	if !reflect.DeepEqual(defaultEmptySquare, emptySquare) {
		t.Fatalf("expected empty: %#v, got a different: %#v", defaultEmptySquare, emptySquare)
	}

	defaultEmptyCircle := cf.GetShape("circle")
	emptyCircle := cf.GetShape("circle")
	emptyCircle.Draw()
	if !reflect.DeepEqual(defaultEmptyCircle, emptyCircle) {
		t.Fatalf("expected empty: %#v, got a different: %#v", defaultEmptyCircle, emptyCircle)
	}
}
