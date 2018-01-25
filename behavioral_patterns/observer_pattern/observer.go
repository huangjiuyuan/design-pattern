package observer

import (
	"fmt"
)

type Observer struct {
	class   string
	subject *Subject
}

func NewOberver(class string, subject *Subject) *Observer {
	o := &Observer{
		class:   class,
		subject: subject,
	}
	o.subject.Attach(o)
	return o
}

func (o *Observer) Update() {
	fmt.Printf("Class %s updates state to %d\n", o.class, o.subject.GetState())
}
