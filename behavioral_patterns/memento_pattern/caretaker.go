package memento

type CareTaker struct {
	MementoList []*Memento
}

func (ct *CareTaker) Add(m *Memento) {
	ct.MementoList = append(ct.MementoList, m)
}

func (ct *CareTaker) Get(index int) *Memento {
	if index < len(ct.MementoList) {
		return ct.MementoList[index]
	}
	return nil
}
