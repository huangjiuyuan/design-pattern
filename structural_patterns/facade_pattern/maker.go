package facade

type ShapeMaker struct {
	rect
	circle
}

func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{
		rect:   rect{"Rect", 20, 10},
		circle: circle{"Circle", 15},
	}
}

func (sm *ShapeMaker) DrawRect() {
	sm.rect.Draw()
}

func (sm *ShapeMaker) DrawCircle() {
	sm.circle.Draw()
}
