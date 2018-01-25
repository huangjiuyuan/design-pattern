package memento

type Memento struct {
	State string
}

func NewMemento(state string) *Memento {
	return &Memento{
		State: state,
	}
}

func (m *Memento) GetState() string {
	return m.State
}
