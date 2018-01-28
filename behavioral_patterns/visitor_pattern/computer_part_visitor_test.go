package visitor

import (
	"testing"
)

func TestComputerPartVisitor(t *testing.T) {
	computer := NewComputer()
	computer.Accept()
}
