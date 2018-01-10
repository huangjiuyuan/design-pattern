package shape

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCache(t *testing.T) {
	sc := &ShapeCache{}
	sc.LoadCache()

	expectedRect := &rect{
		shape{"rect", "1"},
		20,
		15,
	}
	cr := sc.GetShape("1")
	fmt.Printf("ID: %s\t", cr.GetID())
	cr.Draw()
	if !reflect.DeepEqual(expectedRect, cr) {
		t.Fatalf("expected rect: %#v, got a different: %#v", expectedRect, cr)
	}

	expectedSquare := &square{
		shape{"square", "2"},
		15,
	}
	cs := sc.GetShape("2")
	fmt.Printf("ID: %s\t", cs.GetID())
	cs.Draw()
	if !reflect.DeepEqual(expectedSquare, cs) {
		t.Fatalf("expected square: %#v, got a different: %#v", expectedSquare, cs)
	}

	expectedCircle := &circle{
		shape{"circle", "50"},
		10,
	}
	cc := sc.GetShape("3")
	fmt.Printf("ID: %s\t", cc.GetID())
	cc.Draw()
	cc.SetID("50")
	if !reflect.DeepEqual(expectedCircle, cc) {
		t.Fatalf("expected circle: %#v, got a different: %#v", expectedCircle, cc)
	}

	// Only the ID of cloned circle is changed. The ID of circle in the cache is not changed.
	if sc.ShapeMap["3"].GetID() != "3" {
		t.Errorf("expected ID: %#v, got a different: %#v", 3, sc.ShapeMap["3"].GetID())
	}
}
