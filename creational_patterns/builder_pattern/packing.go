package builder

type Packing interface {
	Pack() string
}

type Wrapper struct{}

func (w *Wrapper) Pack() string {
	return "Wrapper"
}

type Bottle struct{}

func (b *Bottle) Pack() string {
	return "Bottle"
}
