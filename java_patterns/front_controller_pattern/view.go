package view

import (
	"fmt"
)

type HomeView struct{}

func (hv *HomeView) Show() {
	fmt.Println("Displaying Home Page")
}

type StudentView struct{}

func (sv *StudentView) Show() {
	fmt.Println("Displaying Student Page")
}
