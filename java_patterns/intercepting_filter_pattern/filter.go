package filter

import (
	"fmt"
)

type Filter interface {
	Execute(req string)
}

type AuthenticationFilter struct{}

func (af *AuthenticationFilter) Execute(req string) {
	fmt.Printf("Authenticating request: %s\n", req)
}

type DebugFilter struct{}

func (df *DebugFilter) Execute(req string) {
	fmt.Printf("Request log: %s\n", req)
}
