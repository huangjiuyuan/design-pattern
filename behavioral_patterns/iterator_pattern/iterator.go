package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Container interface {
	GetIterator() Iterator
}

type NameRepository struct {
	Names []string
}

func (nr *NameRepository) GetIterator() Iterator {
	return &NameIterator{
		NameRepository: *nr,
		Index:          0,
	}
}

type NameIterator struct {
	NameRepository
	Index int
}

func (ni *NameIterator) HasNext() bool {
	return ni.Index < len(ni.Names)
}

func (ni *NameIterator) Next() interface{} {
	if ni.HasNext() {
		name := ni.Names[ni.Index]
		ni.Index++
		return name
	}
	return nil
}
