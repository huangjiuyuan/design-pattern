package interpreter

import (
	"strings"
)

type Expression interface {
	Interpret(context string) bool
}

type TerminalExpression struct {
	Data string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{
		Data: data,
	}
}

func (te *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, te.Data) {
		return true
	}
	return false
}

type OrExpression struct {
	Expr1 Expression
	Expr2 Expression
}

func NewOrExpression(expr1, expr2 Expression) *OrExpression {
	return &OrExpression{
		Expr1: expr1,
		Expr2: expr2,
	}
}

func (te *OrExpression) Interpret(context string) bool {
	return te.Expr1.Interpret(context) || te.Expr2.Interpret(context)
}

type AndExpression struct {
	Expr1 Expression
	Expr2 Expression
}

func NewAndExpression(expr1, expr2 Expression) *AndExpression {
	return &AndExpression{
		Expr1: expr1,
		Expr2: expr2,
	}
}

func (te *AndExpression) Interpret(context string) bool {
	return te.Expr1.Interpret(context) && te.Expr2.Interpret(context)
}
