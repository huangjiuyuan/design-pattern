package shape

import (
	"fmt"
)

type Shape interface {
	Draw()
	GetType() string
	GetID() string
	SetID(string)
	Clone() Shape
}

type shape struct {
	shapeType string
	id        string
}

func (s *shape) GetType() string {
	return s.shapeType
}

func (s *shape) GetID() string {
	return s.id
}

func (s *shape) SetID(id string) {
	s.id = id
}

type rect struct {
	shape
	width, height float64
}

func (r *rect) Draw() {
	fmt.Printf("Drawing a %s with width %f, height %f\n", r.shapeType, r.width, r.height)
}

func (r *rect) Clone() Shape {
	clone := *r
	return &clone
}

type square struct {
	shape
	width float64
}

func (s *square) Draw() {
	fmt.Printf("Drawing a %s with width %f\n", s.shapeType, s.width)
}

func (s *square) Clone() Shape {
	clone := *s
	return &clone
}

type circle struct {
	shape
	radius float64
}

func (c *circle) Draw() {
	fmt.Printf("Drawing a %s with radius %f\n", c.shapeType, c.radius)
}

func (c *circle) Clone() Shape {
	clone := *c
	return &clone
}

type empty struct {
	shape
}

func (e *empty) Draw() {
	fmt.Println("Drawing nothing")
}

func (e *empty) Clone() Shape {
	clone := *e
	return &clone
}
