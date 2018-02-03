package locator

import (
	"fmt"
)

type Service interface {
	Execute()
	GetName() string
}

type ServiceA struct{}

func (a *ServiceA) Execute() {
	fmt.Println("Executing ServiceA")
}

func (a *ServiceA) GetName() string {
	return "ServiceA"
}

type ServiceB struct{}

func (b *ServiceB) Execute() {
	fmt.Println("Executing ServiceB")
}

func (b *ServiceB) GetName() string {
	return "ServiceB"
}
