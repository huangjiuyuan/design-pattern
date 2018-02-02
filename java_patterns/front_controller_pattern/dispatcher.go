package view

import (
	"fmt"
)

type Dispatcher struct {
	studentView *StudentView
	homeView    *HomeView
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		studentView: new(StudentView),
		homeView:    new(HomeView),
	}
}

func (d *Dispatcher) Dispatch(req string) {
	if req == "student" {
		d.studentView.Show()
	} else if req == "home" {
		d.homeView.Show()
	} else {
		fmt.Println("Invalid")
	}
}
