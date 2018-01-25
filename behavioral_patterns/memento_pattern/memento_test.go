package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	originator := new(Originator)
	caretaker := new(CareTaker)

	originator.SetState("State 1")
	originator.SetState("State 2")
	caretaker.Add(originator.SaveStateToMemento())
	originator.SetState("State 3")
	caretaker.Add(originator.SaveStateToMemento())
	originator.SetState("State 4")

	fmt.Printf("Current State: %s\n", originator.GetState())
	if originator.GetState() != "State 4" {
		t.Errorf("expected State 4, got %s", originator.GetState())
	}

	originator.GetStateFromMemento(caretaker.Get(0))
	fmt.Printf("Restore to State: %s\n", originator.GetState())
	if originator.GetState() != "State 2" {
		t.Errorf("expected State 2, got %s", originator.GetState())
	}

	originator.GetStateFromMemento(caretaker.Get(1))
	fmt.Printf("Restore to State: %s\n", originator.GetState())
	if originator.GetState() != "State 3" {
		t.Errorf("expected State 3, got %s", originator.GetState())
	}
}
