package factory

type Factory interface {
	GetColor(c string) Color
	GetShape(s string) Shape
}

type FactoryProducer struct{}

func (fp *FactoryProducer) GetFactory(f string) Factory {
	if f == "shape" {
		return &ShapeFactory{}
	} else if f == "color" {
		return &ColorFactory{}
	}
	return &EmptyFactory{}
}

type ShapeFactory struct{}

func (sf *ShapeFactory) GetColor(c string) Color {
	return &emptyColor{}
}

func (sf *ShapeFactory) GetShape(s string) Shape {
	if s == "rect" {
		return &rect{
			shapeType: "rect",
			width:     20,
			height:    15,
		}
	} else if s == "square" {
		return &square{
			shapeType: "square",
			width:     15,
		}
	} else if s == "circle" {
		return &circle{
			shapeType: "circle",
			radius:    10,
		}
	}
	return &emptyShape{}
}

type ColorFactory struct{}

func (cf *ColorFactory) GetColor(c string) Color {
	if c == "red" {
		return &red{
			color: "red",
			red:   255,
			green: 0,
			blue:  0,
		}
	} else if c == "green" {
		return &green{
			color: "green",
			red:   0,
			green: 255,
			blue:  0,
		}
	} else if c == "blue" {
		return &blue{
			color: "blue",
			red:   0,
			green: 0,
			blue:  255,
		}
	}
	return &emptyColor{}
}

func (cf *ColorFactory) GetShape(s string) Shape {
	return &emptyShape{}
}

type EmptyFactory struct{}

func (ef *EmptyFactory) GetColor(c string) Color {
	return &emptyColor{}
}

func (ef *EmptyFactory) GetShape(s string) Shape {
	return &emptyShape{}
}
