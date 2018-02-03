package student

import (
	"fmt"
	"reflect"
)

type StudentTransferObject struct {
	name   string
	number int
}

func NewStudentTransferObject(name string, number int) *StudentTransferObject {
	return &StudentTransferObject{
		name:   name,
		number: number,
	}
}

func (to *StudentTransferObject) GetName() string {
	return to.name
}

func (to *StudentTransferObject) SetName(name string) {
	to.name = name
}

func (to *StudentTransferObject) GetNumber() int {
	return to.number
}

func (to *StudentTransferObject) SetNumber(number int) {
	to.number = number
}

type StudentBusinessObject struct {
	Students []*StudentTransferObject
}

func NewStudentBusinessObject() *StudentBusinessObject {
	students := make([]*StudentTransferObject, 0)
	s1 := NewStudentTransferObject("Robert", 1)
	s2 := NewStudentTransferObject("John", 2)
	s3 := NewStudentTransferObject("Mary", 3)
	students = append(students, s1, s2, s3)

	return &StudentBusinessObject{
		Students: students,
	}
}

func (bo *StudentBusinessObject) DeleteStudent(to *StudentTransferObject) {
	for i, s := range bo.Students {
		if reflect.DeepEqual(s, to) {
			bo.Students = append(bo.Students[:i], bo.Students[i+1:]...)
			fmt.Printf("Deleting student %d from database\n", to.GetNumber())
		}
	}
}

func (bo *StudentBusinessObject) GetAllStudents() []*StudentTransferObject {
	return bo.Students
}

func (bo *StudentBusinessObject) GetStudent(number int) *StudentTransferObject {
	for _, s := range bo.Students {
		if s.GetNumber() == number {
			return s
		}
	}
	return nil
}

func (bo *StudentBusinessObject) UpdateStudent(to *StudentTransferObject) {
	for i, s := range bo.Students {
		if s.GetNumber() == to.GetNumber() {
			bo.Students[i] = to
			fmt.Printf("Updating student %d in the database\n", s.GetNumber())
		}
	}
}
