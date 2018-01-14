package bridge

import (
	"fmt"
)

// DrawAPI defines the bridge interface.
type DrawAPI interface {
	DrawCircle(radius int, x int, y int)
}

type RedCircle struct {
	Radius int
	X      int
	Y      int
}

func (rc *RedCircle) DrawCircle(radius int, x int, y int) {
	fmt.Printf("Drawing red circle with radius: %d, x: %d, y: %d\n", radius, x, y)
}

type GreenCircle struct {
	Radius int
	X      int
	Y      int
}

func (rc *GreenCircle) DrawCircle(radius int, x int, y int) {
	fmt.Printf("Drawing green circle with radius: %d, x: %d, y: %d\n", radius, x, y)
}
