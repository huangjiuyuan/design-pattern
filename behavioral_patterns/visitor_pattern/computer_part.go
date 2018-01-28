package visitor

type ComputerPart interface {
	Accept()
}

type Keyboard struct {
	ComputerPartVisitor ComputerPartVisitor
}

func (k *Keyboard) Accept() {
	k.ComputerPartVisitor.VisitKeyboard(k)
}

type Monitor struct {
	ComputerPartVisitor ComputerPartVisitor
}

func (m *Monitor) Accept() {
	m.ComputerPartVisitor.VisitMonitor(m)
}

type Mouse struct {
	ComputerPartVisitor ComputerPartVisitor
}

func (m *Mouse) Accept() {
	m.ComputerPartVisitor.VisitMouse(m)
}

type Computer struct {
	ComputerPartVisitor ComputerPartVisitor
	ComputerParts       []ComputerPart
}

func NewComputer() *Computer {
	return &Computer{
		ComputerPartVisitor: &computerPartVisitor{},
		ComputerParts: []ComputerPart{
			&Mouse{ComputerPartVisitor: &computerPartVisitor{}},
			&Keyboard{ComputerPartVisitor: &computerPartVisitor{}},
			&Monitor{ComputerPartVisitor: &computerPartVisitor{}},
		},
	}
}

func (c *Computer) Accept() {
	for _, cp := range c.ComputerParts {
		cp.Accept()
	}
	c.ComputerPartVisitor.VisitComputer(c)
}
