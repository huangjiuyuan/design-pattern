package filter

import (
	"fmt"
)

type Target struct{}

func (t *Target) Execute(req string) {
	fmt.Printf("Executing request: %s\n", req)
}
