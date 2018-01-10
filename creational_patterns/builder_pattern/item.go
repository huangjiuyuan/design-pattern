package builder

type Item interface {
	Name() string
	Packing() Packing
	Price() float64
}

type VegBurger struct{}

func (vb *VegBurger) Name() string {
	return "Veg Burger"
}

func (vb *VegBurger) Packing() Packing {
	return new(Wrapper)
}

func (vb *VegBurger) Price() float64 {
	return 25.0
}

type ChickenBurger struct{}

func (cb *ChickenBurger) Name() string {
	return "Chicken Burger"
}

func (cb *ChickenBurger) Packing() Packing {
	return new(Wrapper)
}

func (cb *ChickenBurger) Price() float64 {
	return 40.0
}

type Coke struct{}

func (c *Coke) Name() string {
	return "Coke"
}

func (c *Coke) Packing() Packing {
	return new(Bottle)
}

func (c *Coke) Price() float64 {
	return 15.0
}

type Pepsi struct{}

func (p *Pepsi) Name() string {
	return "Pepsi"
}

func (p *Pepsi) Packing() Packing {
	return new(Bottle)
}

func (p *Pepsi) Price() float64 {
	return 10.0
}
