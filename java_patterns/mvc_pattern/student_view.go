package mvc

import (
	"fmt"
)

// Student view.
type StudentView struct{}

func (sv *StudentView) PrintStudentDetails(name string, row int) {
	fmt.Printf("Student %s sits at row %d\n", name, row)
}
