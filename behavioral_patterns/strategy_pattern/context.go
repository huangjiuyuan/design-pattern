package strategy

type Context struct {
	Strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{
		Strategy: strategy,
	}
}

func (ctx *Context) ExecuteStrategy(a, b int) int {
	return ctx.Strategy.DoOperation(a, b)
}
