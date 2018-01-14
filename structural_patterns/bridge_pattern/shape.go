package bridge

// Shape contains the bridge interface.
type Shape struct {
	DrawAPI DrawAPI
}

func (s *Shape) Shape(a DrawAPI) {
	s.DrawAPI = a
}

// Circle inherits Shape.
type Circle struct {
	Shape  Shape
	Radius int
	X      int
	Y      int
}

func NewCircle(radius int, x int, y int, d DrawAPI) *Circle {
	c := new(Circle)
	c.Shape.DrawAPI = d
	c.Radius, c.X, c.Y = radius, x, y
	return c
}

func (c *Circle) Draw() {
	c.Shape.DrawAPI.DrawCircle(c.Radius, c.X, c.Y)
}
