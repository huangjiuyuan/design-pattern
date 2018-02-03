package locator

import (
	"fmt"
)

type InitialContext struct{}

func (ic *InitialContext) LookUp(jndi string) interface{} {
	if jndi == "ServiceA" {
		fmt.Println("Looking up and creating a new ServiceA object")
		return new(ServiceA)
	} else if jndi == "ServiceB" {
		fmt.Println("Looking up and creating a new ServiceB object")
		return new(ServiceB)
	}
	return nil
}
