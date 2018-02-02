package composite

type CompositeEntity struct {
	cgo *CoarseGrainedObject
}

func NewCompositeEntity() *CompositeEntity {
	return &CompositeEntity{
		cgo: NewCoarseGrainedObject(),
	}
}

func (ce *CompositeEntity) SetData(a, b string) {
	ce.cgo.SetData(a, b)
}

func (ce *CompositeEntity) GetData() []string {
	return ce.cgo.GetData()
}
