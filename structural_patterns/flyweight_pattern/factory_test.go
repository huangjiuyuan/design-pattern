package flyweight

import (
	"math/rand"
	"testing"
)

func TestShapeFactory(t *testing.T) {
	sf := &ShapeFactory{
		CircleMap: make(map[string]Shape, 0),
	}
	colors := []string{"Green", "Blue", "Yellow", "Red", "Purple"}
	for _, color := range colors {
		circle := sf.GetCircle(color).(*circle)
		circle.SetPosition(rand.Intn(100), rand.Intn(100))
		circle.SetRadius(rand.Float64() * 50)
	}

	for _, circle := range sf.CircleMap {
		circle.Draw()
	}
}
