package student

import (
	"fmt"
	"testing"
)

func TestStudentDao(t *testing.T) {
	var dao StudentDao
	dao = NewStudentDaoImplement()
	for _, s := range dao.GetAllStudents() {
		fmt.Printf("[Student] Name: %s, Row: %d\n", s.GetName(), s.GetRow())
	}

	dao.UpdateStudent(NewStudent(0, "Mary"))
	fmt.Printf("Student at row 0: %v\n", dao.GetStudent(0))
}
