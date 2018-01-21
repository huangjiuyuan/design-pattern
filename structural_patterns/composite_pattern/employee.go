package employee

import (
	"fmt"
)

type Employee struct {
	Name         string
	Dept         string
	Salary       int
	Subordinates []*Employee
}

func NewEmployee(name string, dept string, salary int) *Employee {
	e := new(Employee)
	e.Name = name
	e.Dept = dept
	e.Salary = salary
	e.Subordinates = make([]*Employee, 0)
	return e
}

func (e *Employee) AddSubordinate(subordinate *Employee) {
	e.Subordinates = append(e.Subordinates, subordinate)
}

func (e *Employee) RemoveSubordinate(subordinate *Employee) {
	for i, s := range e.Subordinates {
		if s == e.Subordinates[i] {
			e.Subordinates = append(e.Subordinates[:i], e.Subordinates[i+1:]...)
		}
	}
}

func (e *Employee) GetSubordinate() []*Employee {
	return e.Subordinates
}

func (e *Employee) ToString() {
	fmt.Printf("Employee [Name: %s, Dept: %s, Salary: %d]\n", e.Name, e.Dept, e.Salary)
}
