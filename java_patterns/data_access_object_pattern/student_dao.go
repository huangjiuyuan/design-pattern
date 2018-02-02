package student

import (
	"fmt"
	"reflect"
)

type StudentDao interface {
	GetAllStudents() []*Student
	GetStudent(row int) *Student
	UpdateStudent(s *Student)
	DeleteStudent(s *Student)
}

type StudentDaoImplement struct {
	// Students slice as database
	Students []*Student
}

func NewStudentDaoImplement() *StudentDaoImplement {
	students := make([]*Student, 0)
	students = append(students, &Student{0, "Robert"})
	students = append(students, &Student{1, "John"})
	return &StudentDaoImplement{
		Students: students,
	}
}

func (dao *StudentDaoImplement) DeleteStudent(s *Student) {
	for i, v := range dao.Students {
		if reflect.DeepEqual(s, v) {
			dao.Students = append(dao.Students[:i], dao.Students[i+1:]...)
			fmt.Printf("Student %s has been deleted\n", s.GetName())
		}
	}
}

func (dao *StudentDaoImplement) GetAllStudents() []*Student {
	return dao.Students
}

func (dao *StudentDaoImplement) GetStudent(row int) *Student {
	for _, v := range dao.Students {
		if v.row == row {
			return v
		}
	}
	return nil
}

func (dao *StudentDaoImplement) UpdateStudent(s *Student) {
	for _, v := range dao.Students {
		if v.GetRow() == s.GetRow() {
			v.SetName(s.GetName())
			v.SetRow(s.GetRow())
			fmt.Printf("Student has been updated to %s at %d\n", s.GetName(), s.GetRow())
		}
	}
}
