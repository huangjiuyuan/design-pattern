package shape

type Factory struct{}

func (f *Factory) GetShape(s string) Shape {
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
	return &empty{}
}
