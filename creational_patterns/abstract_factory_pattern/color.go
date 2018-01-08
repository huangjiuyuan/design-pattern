package factory

import (
	"fmt"
)

type Color interface {
	Fill()
}

type red struct {
	color            string
	red, green, blue int16
}

func (r *red) Fill() {
	fmt.Printf("Filling with color %s with RGB %d, %d, %d", r.color, r.red, r.green, r.blue)
}

type green struct {
	color            string
	red, green, blue int16
}

func (g *green) Fill() {
	fmt.Printf("Filling with color %s with RGB %d, %d, %d", g.color, g.red, g.green, g.blue)
}

type blue struct {
	color            string
	red, green, blue int16
}

func (b *blue) Fill() {
	fmt.Printf("Filling with color %s with RGB %d, %d, %d", b.color, b.red, b.green, b.blue)
}

type emptyColor struct{}

func (e *emptyColor) Fill() {
	fmt.Println("Filling nothing")
}
