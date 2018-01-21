package employee

import (
	"testing"
)

func TestEmplotee(t *testing.T) {
	ceo := NewEmployee("John", "Administration", 40000)
	headSales := NewEmployee("Robert", "Head Sales", 30000)
	headMarketing := NewEmployee("Michel", "Head Marketing", 30000)

	clerk1 := NewEmployee("Laura", "Marketing", 10000)
	clerk2 := NewEmployee("Bob", "Marketing", 10000)

	salesExecutive1 := NewEmployee("Richard", "Sales", 10000)
	salesExecutive2 := NewEmployee("Rob", "Sales", 10000)

	ceo.AddSubordinate(headSales)
	ceo.AddSubordinate(headMarketing)

	headSales.AddSubordinate(salesExecutive1)
	headSales.AddSubordinate(salesExecutive2)

	headMarketing.AddSubordinate(clerk1)
	headMarketing.AddSubordinate(clerk2)

	ceo.ToString()
	for _, e := range ceo.GetSubordinate() {
		e.ToString()
		for _, e := range e.GetSubordinate() {
			e.ToString()
		}
	}
}
