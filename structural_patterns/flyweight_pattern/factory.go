package flyweight

import (
	"fmt"
)

type ShapeFactory struct {
	CircleMap map[string]Shape
}

func (sf *ShapeFactory) GetCircle(color string) Shape {
	_, ok := sf.CircleMap[color]
	if !ok {
		fmt.Printf("Creating circle with color %s\n", color)
		sf.CircleMap[color] = NewCircle(color)
	}
	return sf.CircleMap[color]
}
