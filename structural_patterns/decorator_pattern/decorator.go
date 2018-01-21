package decorator

import (
	"fmt"
)

type RedDecorator struct {
	shape Shape
}

func NewShapeDecorator(shape Shape) *RedDecorator {
	return &RedDecorator{shape}
}

func (ed *RedDecorator) Draw() {
	fmt.Printf("Select color red. ")
	ed.shape.Draw()
}
