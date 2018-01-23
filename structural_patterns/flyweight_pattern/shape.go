package flyweight

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type circle struct {
	shapeType string
	x, y      int
	radius    float64
	color     string
}

func NewCircle(color string) *circle {
	c := new(circle)
	c.color = color
	return c
}

func (c *circle) SetPosition(x int, y int) {
	c.x = x
	c.y = y
}

func (c *circle) SetRadius(radius float64) {
	c.radius = radius
}

func (c *circle) Draw() {
	fmt.Printf("Drawing a %s at (%d, %d) with radius %f, color %s\n", c.shapeType, c.x, c.y, c.radius, c.color)
}

type emptyShape struct{}

func (e *emptyShape) Draw() {
	fmt.Println("Drawing nothing")
}
