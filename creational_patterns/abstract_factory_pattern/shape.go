package factory

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type rect struct {
	shapeType     string
	width, height float64
}

func (r *rect) Draw() {
	fmt.Printf("Drawing a %s with width %f, height %f\n", r.shapeType, r.width, r.height)
}

type square struct {
	shapeType string
	width     float64
}

func (s *square) Draw() {
	fmt.Printf("Drawing a %s with width %f\n", s.shapeType, s.width)
}

type circle struct {
	shapeType string
	radius    float64
}

func (c *circle) Draw() {
	fmt.Printf("Drawing a %s with radius %f\n", c.shapeType, c.radius)
}

type emptyShape struct{}

func (e *emptyShape) Draw() {
	fmt.Println("Drawing nothing")
}
