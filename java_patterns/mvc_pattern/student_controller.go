package mvc

// Student Controller.
type StudentController struct {
	model *Student
	view  *StudentView
}

func NewStudentController(model *Student, view *StudentView) *StudentController {
	return &StudentController{
		model: model,
		view:  view,
	}
}

func (sc *StudentController) SetStudentName(name string) {
	sc.model.SetName(name)
}

func (sc *StudentController) GetStudentName() string {
	return sc.model.GetName()
}

func (sc *StudentController) SetStudentRow(row int) {
	sc.model.SetRow(row)
}

func (sc *StudentController) GetStudentRow() int {
	return sc.model.GetRow()
}

func (sc *StudentController) UpdateView() {
	sc.view.PrintStudentDetails(sc.GetStudentName(), sc.GetStudentRow())
}
