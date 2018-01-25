package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	subject := new(Subject)
	NewOberver("A", subject)
	NewOberver("B", subject)
	NewOberver("C", subject)

	subject.SetState(15)
	subject.SetState(-3)
	subject.SetState(80)
}
