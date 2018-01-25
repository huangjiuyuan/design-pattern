package memento

type Originator struct {
	State string
}

func (o *Originator) SetState(state string) {
	o.State = state
}

func (o *Originator) GetState() string {
	return o.State
}

func (o *Originator) SaveStateToMemento() *Memento {
	return NewMemento(o.State)
}

func (o *Originator) GetStateFromMemento(m *Memento) {
	o.State = m.State
}
