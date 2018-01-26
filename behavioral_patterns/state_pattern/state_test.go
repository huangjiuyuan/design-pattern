package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	ctx := &Context{}
	start := &StartState{}
	stop := &StopState{}

	start.DoAction(ctx)
	fmt.Println(ctx.GetState())
	stop.DoAction(ctx)
	fmt.Println(ctx.GetState())
}
