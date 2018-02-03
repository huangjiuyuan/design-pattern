package student

import (
	"fmt"
	"testing"
)

func TestStudent(t *testing.T) {
	bo := NewStudentBusinessObject()
	for _, s := range bo.GetAllStudents() {
		fmt.Printf("[Student] %d: %s\n", s.GetNumber(), s.GetName())
	}

	s := NewStudentTransferObject("Michael", 1)
	bo.UpdateStudent(s)
	s = bo.GetStudent(1)
	fmt.Printf("Student 1 is %s\n", s.GetName())
}
