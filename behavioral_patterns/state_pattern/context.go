package state

type Context struct {
	State string
}

func (c *Context) SetState(state string) {
	c.State = state
}

func (c *Context) GetState() string {
	return c.State
}
