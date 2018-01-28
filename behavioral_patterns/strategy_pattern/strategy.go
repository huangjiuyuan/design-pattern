package strategy

type Strategy interface {
	DoOperation(a, b int) int
}

type OperationAdd struct{}

func (oa *OperationAdd) DoOperation(a, b int) int {
	return a + b
}

type OperationSubstract struct{}

func (oa *OperationSubstract) DoOperation(a, b int) int {
	return a - b
}

type OperationMultiply struct{}

func (oa *OperationMultiply) DoOperation(a, b int) int {
	return a * b
}
