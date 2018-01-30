package mvc

import (
	"testing"
)

func TestStudentController(t *testing.T) {
	model := &Student{}
	model.SetName("John")
	model.SetRow(10)

	view := &StudentView{}
	view.PrintStudentDetails(model.GetName(), model.GetRow())

	controller := &StudentController{
		model: model,
		view:  view,
	}
	controller.SetStudentName("Lee")
	controller.SetStudentRow(25)
	controller.UpdateView()
}
