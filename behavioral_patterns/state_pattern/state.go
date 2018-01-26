package state

import (
	"fmt"
)

type State interface {
	DoAction(context Context)
}

type StartState struct{}

func (ss *StartState) DoAction(ctx *Context) {
	fmt.Println("Player is in start state")
	ctx.SetState("Start")
}

type StopState struct{}

func (ss *StopState) DoAction(ctx *Context) {
	fmt.Println("Player is in stop state")
	ctx.SetState("Stop")
}
