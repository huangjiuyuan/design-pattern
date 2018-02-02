package composite

type DependentObjectA struct {
	data string
}

func (do *DependentObjectA) SetData(data string) {
	do.data = data
}

func (do *DependentObjectA) GetData() string {
	return do.data
}

type DependentObjectB struct {
	data string
}

func (do *DependentObjectB) SetData(data string) {
	do.data = data
}

func (do *DependentObjectB) GetData() string {
	return do.data
}
