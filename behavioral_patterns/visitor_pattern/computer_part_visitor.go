package visitor

import (
	"fmt"
)

type ComputerPartVisitor interface {
	VisitComputer(computer *Computer)
	VisitMouse(mouse *Mouse)
	VisitKeyboard(keyboard *Keyboard)
	VisitMonitor(monitor *Monitor)
}

type computerPartVisitor struct{}

func (cpv *computerPartVisitor) VisitComputer(computer *Computer) {
	fmt.Println("Visiting computer")
}

func (cpv *computerPartVisitor) VisitMouse(mouse *Mouse) {
	fmt.Println("Visiting mouse")
}

func (cpv *computerPartVisitor) VisitKeyboard(keyboard *Keyboard) {
	fmt.Println("Visiting keyboard")
}

func (cpv *computerPartVisitor) VisitMonitor(monitor *Monitor) {
	fmt.Println("Visiting monitor")
}
