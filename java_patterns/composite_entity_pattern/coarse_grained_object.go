package composite

type CoarseGrainedObject struct {
	objectA *DependentObjectA
	objectB *DependentObjectB
}

func NewCoarseGrainedObject() *CoarseGrainedObject {
	return &CoarseGrainedObject{
		objectA: new(DependentObjectA),
		objectB: new(DependentObjectB),
	}
}

func (cgo *CoarseGrainedObject) SetData(a, b string) {
	cgo.objectA.SetData(a)
	cgo.objectB.SetData(b)
}

func (cgo *CoarseGrainedObject) GetData() []string {
	return []string{cgo.objectA.GetData(), cgo.objectB.GetData()}
}
